package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func handleConnectionError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// open connection with RabbitMQ
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleConnectionError(err, "Failed to establish connection")
	defer connection.Close()

	// open channel
	channel, err := connection.Channel()
	handleConnectionError(err, "Failed to open channel")
	defer channel.Close()
}
