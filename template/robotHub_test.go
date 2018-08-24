package template

import (
	"testing"
	"time"
	. "tkBot-template-go/strategy"
)

func TestNewRobotHub(t *testing.T) {
	robotHub:=NewRobotHub(RobotHubName,WebsocketServer,Heartbeat_Interval)
	go robotHub.Start()
	para:=&Parameters{
		"test robot",
		5,
	}
	kill_chan:=make(chan byte)
	go robotHub.StartRobot(para,kill_chan)
	time.Sleep(40*time.Second)
	kill_chan<-0
	time.Sleep(20*time.Second)
}
