package main

import (

	"encoding/json"
	"fmt"
	"time"
	"tkBot-template-go/template"
)




func main(){
	robotHub:=template.NewRobotHub(template.RobotHubName,template.WebsocketServer,template.Heartbeat_Interval)
	robotHub.Start()
	for  {
		time.Sleep(24*time.Hour)
	}
}
func test(){
	robot_list:=make([]string,0)
	robot_list=append(robot_list, "sss")
	robot_list=append(robot_list, "vvv")
	robot_list=append(robot_list, "lll")



	data,_:=json.Marshal(&robot_list)
	fmt.Println(string(data))

	var robot_list2 []string
	cha1:=make(chan byte)
	cha2:=make(chan byte)
	cha3:=make(chan byte)
	cha4:=make(chan byte)
	var testMap=map[string] chan byte{
		"sss":cha1,
		"vvv":cha2,
		"lll":cha3,

	}
	go func() {
		for {
			select {
			case <-cha1:
				fmt.Println("cha1")
			case <-cha2:
				fmt.Println("cha2")
			case <-cha3:
				fmt.Println("cha3")
			case <-cha4:
				fmt.Println("cha4")
			}
		}
	}()

	json.Unmarshal(data,&robot_list2)
	fmt.Println(testMap)
	robots := make([]string, 0, len(testMap))
	for k := range testMap {
		robots = append(robots, k)
	}
	fmt.Println(robots)
	for _,robot:=range robot_list2{
		testMap[robot]<-0;
		delete(testMap, robot)
	}
	fmt.Println(testMap)
	time.Sleep(2*time.Second)
	robots = make([]string, 0, len(testMap))
	for k := range testMap {
		robots = append(robots, k)
	}
	fmt.Println(robots)
}