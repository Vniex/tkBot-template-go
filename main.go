package main

import (

	"tkBot-template-go/template"
)




func main(){
	robotHub:=template.NewRobotHub(template.RobotHubName,template.WebsocketServer,template.Heartbeat_Interval)
	robotHub.Start()
}