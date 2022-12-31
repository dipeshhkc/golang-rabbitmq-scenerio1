package main

import (
	"fmt"
	"rabbitmq-producer-site-check/constants"
	"rabbitmq-producer-site-check/lib"
	"rabbitmq-producer-site-check/router"
	"rabbitmq-producer-site-check/utils"
)

func main() {

	connection, channel := lib.SetupRabbbitMQConnectionChannel()
	defer connection.Close()
	defer channel.Close()

	requestQueue, err := channel.QueueDeclare(
		constants.QUEUE, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)

	utils.FailOnError(err, "Failed to register a queue")

	request, err := channel.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		false,             // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	utils.FailOnError(err, "Failed to register a listener in queue")

	go func() {
		for d := range request {
			switch d.Type {
			case constants.TASK1:
				fmt.Println("IN PROGRESS - TASK1")
				//1. Perform Task1 Logic
				//2. [IF NEEEDED] Store results of Task1 to DB
				//3. [IF NEEEDED] Send results to the frontend using Websocket
				//4. remove msg from queue
				d.Ack(false)
			case constants.TASK2:
				fmt.Println("IN PROGRESS - TASK2")
				//1. Perform Task2 Logic
				//2. [IF NEEEDED] Store results of Task2 to DB
				//3. [IF NEEEDED] Send results to the frontend using Websocket
				//4. remove msg from queue
				d.Ack(false)
			}
		}

	}()

	router.SetupRoutes()

}
