package strategy

import (

	"time"
	"fmt"
	"log"
)

const RobotHubName="托管者测试" //策略托管者名称
const ServerIP="127.0.0.1"   //服务器ip
const ServerPort="8888"		//服务器端口
const ServerURL="/api/v1/ws" //服务器ws路径
const WebsocketServer = "ws://"+ServerIP+":"+ServerPort+ServerURL  //组合URL路径
const Heartbeat_Interval=2   //心跳包间隔

type Robot struct {
	Para *Parameters
}

//策略自定义参数
type Parameters struct{
	RobotName string `json:"robot_name"`  //用以识别机器人，不可修改
	Interval int `json:"interval"`
}

//策略默认参数
func NewDefaultParameters() *Parameters{
	return &Parameters{
		"robot",
		5,
	}
}

func NewRobot(parameters *Parameters)*Robot{
	return &Robot{parameters}
}


//自定义策略
func (r *Robot)Start(){
	r.onTick()

}


func (r *Robot)onTick(){
	for {

		fmt.Printf("%v in tick\n",r.Para.RobotName)
		log.Println(r.Para.Interval)
		time.Sleep(5*time.Second)
	}

}