package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"panda-poc-consumer-fila/config"
	resource "panda-poc-consumer-fila/resources"
	"panda-poc-consumer-fila/server"
	"syscall"

	"github.com/streadway/amqp"
)

func main() {
	amqpConns := make([]*amqp.Connection, config.NumberOfConnections)
	var err error
	for i := 0; i < config.NumberOfConnections; i++ {
		amqpConns[i], err = resource.NewRabbitMQConn()

		if err != nil {
			log.Fatal(err)
		}
		defer amqpConns[i].Close()
	}

	for i := 0; i < config.NumberOfConnections; i++ {
		s := server.NewServer(amqpConns)
		go s.Run()
	}

	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Println("signal.Notify: ", v)
	case done := <-ctx.Done():
		log.Println("ctx.Done: ", done)
	}

	if err != nil {
		log.Println("Startando Consumer: ", err)
		cancel()
	}
}
