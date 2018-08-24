package template

import (
	"encoding/json"
	"log"
)

const(
	CmdType_HEARTBEAT = iota+1
	CmdType_KILL
	CmdType_PUSHDATA
	CmdType_START

)


type RobotHubMsg struct {
	RobotHubName string `json:"robot_hub_name"`
	Cmd   int      `json:"cmd"`
	Data  string `json:"data"`
}


func NewRobotHubMsg(robot_hub_name string,cmd int,data string) *RobotHubMsg{
	return &RobotHubMsg{robot_hub_name,cmd,data}
}

func (r *RobotHubMsg)ToBytes() ([]byte,error){
	msg, err := json.Marshal(r)
	if err != nil {
		log.Printf("Fail to package robotMsg :%v", err)
		return nil,err
	}
	return msg,nil
}



func ParseRobotMsg(message []byte) *RobotHubMsg {
	var data RobotHubMsg
	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Println("Fail to parse message:%v", err)
		return nil
	}
	return &data
}

func PackageRobotMsg(robot_name string,cmd int,   data string) []byte {
	var req = RobotHubMsg{
		RobotHubName:robot_name,
		Cmd:   cmd,
		Data:  data,
	}

	msg, err := json.Marshal(req)
	if err != nil {
		log.Println("Fail to packageResponseMsg:%v", err)
		return nil
	}
	return msg
}
