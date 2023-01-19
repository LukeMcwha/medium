package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

// constant created for NATs reply timeout.
const timeout = 3 * time.Second

func main() {
	// Connect to the NATs server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Println("Connection Err:> ", err)
		return
	}

	// Lets send a message!
	fmt.Println("Sending message to subject...")

	msg, err := nc.Request("hello/world", []byte("Hi there other service! :D"), timeout)
	if err != nil {
		fmt.Println("Error while sending message:> ", err)
		fmt.Println("Shutting down process due to error...")
		return
	}

	// Let's just print out the response for the message.
	// Since we know the response data is a string, we will set it to a string.
	fmt.Println("Message Response:> ", string(msg.Data))

	// Nothing left to do, shutdown the process.
	fmt.Println("All processes are complete, time to shutdown...")
}
