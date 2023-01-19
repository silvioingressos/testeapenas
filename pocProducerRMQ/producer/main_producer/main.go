package main

import (
	"log"
	resource "panda-poc-fila/resources"
	"panda-poc-fila/server"
)

func main() {

	amqpConn, err := resource.NewRabbitMQConn()

	if err != nil {
		log.Fatal(err)
	}
	defer amqpConn.Close()

	s := server.NewServer(amqpConn)

	s.Run()
}
