package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"panda-poc-fila/producer"
	"strconv"
	"syscall"

	"github.com/streadway/amqp"
)

type Server struct {
	amqpConn *amqp.Connection
	producer producer.PublisherIF
}

func NewServer(amqpConn *amqp.Connection) *Server {
	return &Server{amqpConn: amqpConn}
}

func (s *Server) Run() error {

	producer, err := producer.NewPublisher()
	if err != nil {
		log.Panic("Olha isso: ", err)
	}

	if err != nil {
		log.Panic("Olha isso: ", err)
	}

	ctx, _ := context.WithCancel(context.Background())
	// for i := 0; i < 100000; i++ {
	for i := 0; i < 2; i++ {
		algo := "Eu acho q vi uma mensagem... " + strconv.Itoa(i)

		go func(algo string) {
			err := producer.PublishToQueue(algo)
			if err != nil {
				log.Println("Geralmente funciona quando faz certo... :", err)
			}
		}(algo)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Println("signal.Notify: ", v)
	case done := <-ctx.Done():
		log.Println("ctx.Done: ", done)
	}

	return nil
}
