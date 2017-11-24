package main

import (
	"fmt"
	"mylib/Rmqtt"
	"os"
)

func main() {

	var mqtt Rmqtt.T_MQTT
	if re := mqtt.CreateClient("tcp://192.168.0.10:1883", "simple"); !re {
		fmt.Printf("create client error!\n")
		os.Exit(1)
	}

}
