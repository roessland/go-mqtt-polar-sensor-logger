package main

import (
	"github.com/fhmq/hmq/broker"
	"log"
)

func serveMqttBroker() {
	config, err := broker.LoadConfig("hmq.config")
	if err != nil {
		log.Fatal("configure broker config error: ", err)
	}

	b, err := broker.NewBroker(config)
	if err != nil {
		log.Fatal("New Broker error: ", err)
	}
	b.Start()
}
