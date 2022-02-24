package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("gordic")
	opts.SetOrderMatters(true)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.SubscribeMultiple(map[string]byte{
		"psl/hr":  0,
		"psl/ecg": 0,
	}, onMessage); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(60 * time.Second)

	if token := c.Unsubscribe("psl/hr"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}

func onMessage(c mqtt.Client, m mqtt.Message) {
	fmt.Println(m.Topic(), string(m.Payload()))
}
