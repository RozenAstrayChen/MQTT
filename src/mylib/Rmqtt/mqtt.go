package Rmqtt

import (
	"fmt"
	//"os"
	"time"
	//import the Paho Go MQTT library

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type T_MQTT struct {
	//define func for the default message handler
	opts     *MQTT.ClientOptions
	callback MQTT.MessageHandler
	client   MQTT.Client
	token    MQTT.Token
}

//define func for the default message handler
//var f MQTT.MessageHandler =
func (t T_MQTT) SimpleCallback(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
func (t T_MQTT) CreateClient(url string, id string) bool {
	fmt.Printf("create start!\n")
	t.opts = MQTT.NewClientOptions().AddBroker(url)
	t.opts.SetClientID(id)
	t.callback = t.SimpleCallback
	t.opts.SetDefaultPublishHandler(t.callback)
	//create and start client using the above ClientOptions
	t.client = MQTT.NewClient(t.opts)
	if t.token = t.client.Connect(); t.token.Wait() && t.token.Error() != nil {
		panic(t.token.Error())
		return false
	}
	//test psuh
	t.Publish("mqtt-go", "test on!", 3)
	return true
}

//publish mesage to broker
func (t T_MQTT) Publish(toipc string, msg string, times int) {
	for i := 0; i < times; i++ {
		t.token = t.client.Publish(toipc, 0, false, msg)
	}
	//need to delay because process will shutdown when clint pushing
	time.Sleep(3 * time.Second)
}
func (t T_MQTT) Disconnect() {
	t.client.Disconnect(250)
}
func (t T_MQTT) TestType() {
	fmt.Printf("test1\n")
}
