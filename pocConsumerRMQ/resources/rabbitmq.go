package resource

import (
	"fmt"
	"panda-poc-consumer-fila/config"

	"github.com/streadway/amqp"
)

func NewRabbitMQConn() (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
	return amqp.Dial(connAddr)
}
