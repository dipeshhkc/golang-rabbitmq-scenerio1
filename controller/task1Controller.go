package controller

import (
	"net/http"
	"rabbitmq-producer-site-check/constants"
	"rabbitmq-producer-site-check/lib"

	"rabbitmq-producer-site-check/utils"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func Task1Controller(c *gin.Context) {
	ch := lib.RabbitChannel
	err := ch.Publish(
		"",              // exchange
		constants.QUEUE, // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        nil,
			Type:        constants.TASK1,
		})

	utils.FailOnError(err, "Failed to publish a message")

	c.JSON(http.StatusOK, gin.H{
		"message": "Task-1 Received Successfully",
	})
}
