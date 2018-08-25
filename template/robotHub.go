package template

import (
	"log"
	"github.com/gorilla/websocket"
	"encoding/json"
	"time"
	"fmt"
	"tkBot-template-go/strategy"
	"os/exec"
	"io"
	"bufio"
)




type RobotHub struct {
	RobotHubName string
	websocketServer string
	heartBeatInterval int
	wsConn *WsConnection
	robots map[string] chan byte
	Para *strategy.Parameters
}



func NewRobotHub(name,websocketServer string,interval int) *RobotHub{
	conn, _, err := websocket.DefaultDialer.Dial(websocketServer, nil)
	if err != nil {
		log.Printf("Fail to dial: %v", err)
		return nil
	}
	log.Println("RobotHub connect success!!!")
	return &RobotHub{name,
	websocketServer,
	interval,
	NewWsConnection(conn),
	make(map[string] chan byte),
	&strategy.Parameters{}}
}



func (r *RobotHub)ProcMsg(msg *RobotHubMsg){
	// Kill some robots
	if msg.Cmd==CmdType_Kill{
		var robot_list []string
		data:=msg.Data
		err:=json.Unmarshal([]byte(data),&robot_list)
		if err!=nil{
			log.Println(err)
			//r.wsConn.WsWrite(NewRobotMsg())
		}
		for _,robot :=range robot_list{
			r.robots[robot]<-0;
			delete(r.robots, robot)
		}
		//Start single robots
	}else if msg.Cmd==CmdType_Start{
		var para *strategy.Parameters
		data:=msg.Data
		err:=json.Unmarshal([]byte(data),&para)
		if err!=nil{
			log.Println(err)
			//r.wsConn.WsWrite(NewRobotMsg())
		}
		robot_kill_chan:=make(chan byte)
		r.robots[para.RobotName]=robot_kill_chan
		go r.StartRobot(para,robot_kill_chan)

	}

}

func (r *RobotHub) HeartBeat() {
	for{
		time.Sleep(time.Duration(r.heartBeatInterval)*time.Second)
		robots := make([]string, 0, len(r.robots))
		for k := range r.robots {
			robots = append(robots, k)
		}
		data,_:=json.Marshal(&robots)

		msg:=NewRobotHubMsg(r.RobotHubName,CmdType_HeartBeat,string(data))
		if err := r.wsConn.WsWrite(msg); err != nil {
			fmt.Println("heartbeat fail")
			r.wsConn.WsClose()
			break
		}
	}


}

func (r *RobotHub)Register(){
	var msg RobotHubMsg
	para:=strategy.NewDefaultParameters()
	msg.RobotHubName=r.RobotHubName
	msg.Cmd=CmdType_Register
	data,_:=json.Marshal(&para)
	msg.Data=string(data)
	r.wsConn.WsWrite(&msg)
}

func (r *RobotHub)ProcRobotStdout(stdout io.ReadCloser){
	in := bufio.NewScanner(stdout)
	for in.Scan(){
		log.Println(in.Text())
	}

}


func (r *RobotHub)ProcRobotStderr(stderr io.ReadCloser){
	in := bufio.NewScanner(stderr)
	for in.Scan(){
		log.Println(in.Text())
	}

}
func (r *RobotHub)StartRobot(parameters *strategy.Parameters,kill_chan chan byte){
	if r.robots[parameters.RobotName]!=nil{ //todo 已运行同名robot
		return
	}
	r.robots[parameters.RobotName]=make(chan byte)
	log.Println("start robot",parameters)

	p,_:=json.Marshal(parameters)

	cmd := exec.Command("../main", "-mode=robot","-para="+string(p))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("get cmd stdout failed with %s\n", err)
		return
	}
	stderr,err:=cmd.StderrPipe()
	if err != nil {
		log.Fatalf("get cmd stderr failed with %s\n", err)
		return
	}
	 if err := cmd.Start();err!=nil{
		 log.Fatalf("cmd start failed with %s\n", err)
		 return
	 }


	go r.ProcRobotStdout(stdout)
	go r.ProcRobotStderr(stderr)




	for ; ;  {
		select {
		case <-kill_chan:
			fmt.Println("kill robot",r.Para.RobotName)
			if err:=cmd.Process.Kill();err!=nil{
				log.Println(err)
			}
			delete(r.robots, parameters.RobotName)
			return

		}

	}

}
func (r *RobotHub)Start(){
	r.Register()
	go r.wsConn.ProcLoop(func(msg *RobotHubMsg) {
		log.Printf("receive %v\n",msg)
	})
	go r.wsConn.WsReadLoop()
	go r.wsConn.WsWriteLoop()
	r.HeartBeat()
}






