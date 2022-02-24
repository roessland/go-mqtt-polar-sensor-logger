package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	go serveMqttBroker()

	s := waitForSignal()
	log.Println("signal received, exiting broker", s)
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
