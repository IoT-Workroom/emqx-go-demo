package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
)

const (
	serverAddr = "tcp://localhost:1883"
	clientID   = "send_events_client"
)

func pubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker(serverAddr).SetClientID(clientID)

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(pubHandler)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Publish messages
	for i := 0; i < 100; i++ {
		topicName := fmt.Sprintf("testtopic/%d", i)

		token := c.Publish(topicName, 0, false, "Hello World")
		token.Wait()

		time.Sleep(500 * time.Millisecond)
	}

	// Unsubscribe
	if token := c.Unsubscribe("testtopic/#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Disconnect
	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}
