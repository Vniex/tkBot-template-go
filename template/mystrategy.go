package template

import (
	"log"
	"time"
)

const RobotHubName="托管者测试" //策略托管者名称
const ServerIP="127.0.0.1"   //服务器ip
const ServerPort="8888"		//服务器端口
const ServerURL="" //服务器ws路径
const WebsocketServer = "ws://"+ServerIP+":"+ServerPort+ServerURL  //组合URL路径
const Heartbeat_Interval=2   //心跳包间隔

//const WebsocketServer="ws://123.207.167.163:9010/ajaxchattest"

//策略自定义参数
type Parameters struct{
	RobotName string `json:"robot_name"`  //用以识别机器人，不可修改


}

func (r *RobotHub)StartRobot(parameters *Parameters,kill_chan chan byte){
	log.Println(parameters)
	r.Para=parameters


}

func (r *RobotHub)OnTick(){
	for {
		log.Println("robot in tick")
		time.Sleep(5*time.Second)
	}

}