package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
)

func main() {
	// Connect to NATs system
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Println("Connection Err:> ", err)
		return
	}

	// Create a subscription to a subject
	_, err = nc.Subscribe("hello/world", func(msg *nats.Msg) {
		fmt.Println("Message received:> ", string(msg.Data))
		msg.Respond([]byte("Hello There!"))
	})

	// We block here until the connection is closed (in our case just ctrl+c to close the app).
	fmt.Println("Start up complete... Now listening for messages")

	// --- Additional logic added for process to block until we want to close it ---
	// Create channel to listen to system interrupts
	closeAppChannel := make(chan os.Signal, 1)
	signal.Notify(closeAppChannel, os.Interrupt)

	// Block until information sent to the channel.
	<-closeAppChannel
	// ---

	fmt.Println("All processes are complete, time to shutdown...")
}
