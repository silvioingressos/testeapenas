package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"panda-poc-consumer-fila/config"
	"panda-poc-consumer-fila/consumer"

	// "panda-poc-consumer-fila/producer"

	"github.com/streadway/amqp"
)

type Server struct {
	amqpConns []*amqp.Connection
}

func NewServer(amqpConns_paran []*amqp.Connection) *Server {
	return &Server{
		amqpConns: amqpConns_paran,
	}
}

func (s *Server) Run() error {

	log.Println("Instancia do server: ")

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < (config.NumberOfConnections - 1); i++ {

		go func() {
			consumer := consumer.NewImagesConsumer(s.amqpConns[i])
			err := consumer.StartConsumer(
				config.WorkerPoolSize,
				config.ProductionOrderExchange,
				config.CreateQueue,
				config.RoutingKey,
				config.ConsumerTag,
			)
			if err != nil {
				log.Println("StartConsumer: ", err)
				cancel()
			}
		}()
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
