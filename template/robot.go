package template

import (
	"log"
	"github.com/gorilla/websocket"
	"encoding/json"
	"time"
	"fmt"
)




type RobotHub struct {
	RobotHubName string
	websocketServer string
	heartBeatInterval int
	wsConn *WsConnection
	closeRobot map[string] chan byte
	Para *Parameters
}



func NewRobotHub(name,websocketServer string,interval int) *RobotHub{
	conn, _, err := websocket.DefaultDialer.Dial(websocketServer, nil)
	if err != nil {
		log.Printf("Fail to dial: %v", err)
		return nil
	}
	log.Println("RobotHub connect success!!!")
	return &RobotHub{name,websocketServer,interval,NewWsConnection(conn),make(map[string] chan byte),&Parameters{}}
}


func (r *RobotHub)ProcMsg(msg *RobotHubMsg){
	// Kill some robots
	if msg.Cmd==CmdType_KILL{
		var robot_list []string
		data:=msg.Data
		err:=json.Unmarshal([]byte(data),&robot_list)
		if err!=nil{
			log.Println(err)
			//r.wsConn.WsWrite(NewRobotMsg())
		}
		for _,robot :=range robot_list{
			r.closeRobot[robot]<-0;
			delete(r.closeRobot, robot)
		}
		//Start single robots
	}else if msg.Cmd==CmdType_START{
		var para *Parameters
		data:=msg.Data
		err:=json.Unmarshal([]byte(data),&para)
		if err!=nil{
			log.Println(err)
			//r.wsConn.WsWrite(NewRobotMsg())
		}
		robot_kill_chan:=make(chan byte)
		r.closeRobot[para.RobotName]=robot_kill_chan
		go r.StartRobot(para,robot_kill_chan)

	}

}

func (r *RobotHub) HeartBeat() {
	for{
		time.Sleep(time.Duration(r.heartBeatInterval)*time.Second)
		robots := make([]string, 0, len(r.closeRobot))
		for k := range r.closeRobot {
			robots = append(robots, k)
		}
		data,_:=json.Marshal(&robots)

		msg:=NewRobotHubMsg(r.RobotHubName,CmdType_HEARTBEAT,string(data))
		if err := r.wsConn.WsWrite(msg); err != nil {
			fmt.Println("heartbeat fail")
			r.wsConn.WsClose()
			break
		}
	}


}

func (r *RobotHub)Start(){
	go r.wsConn.ProcLoop(func(msg *RobotHubMsg) {
		log.Printf("receive %v\n",msg)
	})
	go r.wsConn.WsReadLoop()
	go r.wsConn.WsWriteLoop()
	r.HeartBeat()
}






