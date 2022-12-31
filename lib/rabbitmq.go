package lib

import (
	"fmt"
	"rabbitmq-producer-site-check/constants"
	"rabbitmq-producer-site-check/utils"

	"github.com/streadway/amqp"
)

var RabbitChannel *amqp.Channel
var rabbitConn *amqp.Connection

//SetupRabbbitMQConnectionChannel -> setup rabbit mq channel
func SetupRabbbitMQConnectionChannel() (*amqp.Connection, *amqp.Channel) {

	//dial
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", constants.USERNAME, constants.PASSWORD, constants.HOST, constants.PORT)

	conn, err := amqp.Dial(url)

	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()

	utils.FailOnError(err, "Failed to open a channel")

	RabbitChannel = ch

	return rabbitConn, RabbitChannel

}
