package main

import (

	"tkBot-template-go/strategy"
	"flag"
	"log"
	"encoding/json"
	"tkBot-template-go/template"
)




func main(){

	mode := flag.String("mode", "robotHub", "start robotHub or robot, robotHub/robot")
	parameter := flag.String("para", "{}", "robot parameterm json string")
	flag.Parse()

	var para strategy.Parameters
	if *mode=="robotHub"{
		robotHub:=template.NewRobotHub(strategy.RobotHubName,strategy.WebsocketServer,strategy.Heartbeat_Interval)
		robotHub.Start()
	}else if *mode=="robot"{
		json.Unmarshal([]byte(*parameter),&para)
		log.Println(*parameter)
		log.Println(para)
		robot:=strategy.NewRobot(&para)
		robot.Start()

	}else{
		panic("no this mode")
	}

}

