package producer

import (
	"encoding/json"
	"log"
	"panda-poc-fila/config"
	"panda-poc-fila/model"
	rabbit "panda-poc-fila/rabbit"
	resource "panda-poc-fila/resources"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

var (
	contador = 0
)

type PublisherIF interface {
	PublishToQueue() error
}

type MessagePublisher struct {
	amqpChan *amqp.Channel
}

func NewPublisher() (*MessagePublisher, error) {
	mqConn, err := resource.NewRabbitMQConn()
	if err != nil {
		return nil, err
	}
	amqpChan, err := mqConn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "p.amqpConn.Channel")
	}
	SetupExchangeAndQueue(amqpChan)
	return &MessagePublisher{amqpChan: amqpChan}, nil
}

func SetupExchangeAndQueue(ch *amqp.Channel) error {

	exchangeArgs := make(amqp.Table)
	exchangeArgs["x-delayed-type"] = "direct"

	rabbit.DeclareExchange(ch, config.ProductionOrderExchange, config.ExchangeKind, exchangeArgs)
	rabbit.DeclareExchange(ch, config.DeadLetterExchange, "direct", nil)
	rabbit.DeclareExchange(ch, config.HospitalExchange, "direct", nil)

	queueNormalArgs := make(amqp.Table)
	queueNormalArgs["x-dead-letter-exchange"] = config.DeadLetterExchange

	queueDLXArgs := make(amqp.Table)
	queueDLXArgs["x-dead-letter-exchange"] = config.ProductionOrderExchange
	queueDLXArgs["x-message-ttl"] = 15000

	rabbit.DeclareQueue(ch, config.CreateQueue, queueNormalArgs)
	rabbit.DeclareQueue(ch, config.DeadLetterQueue, queueDLXArgs)
	rabbit.DeclareQueue(ch, config.HospitalQueue, nil)

	rabbit.BindQueue(ch, config.ProductionOrderExchange, config.CreateQueue, config.RoutingKey)
	rabbit.BindQueue(ch, config.DeadLetterExchange, config.DeadLetterQueue, config.RoutingKey)
	rabbit.BindQueue(ch, config.HospitalExchange, config.HospitalQueue, config.RoutingKey)

	log.Println("Message sent to queue")
	return nil
}

func (p *MessagePublisher) CloseChan() {
	if err := p.amqpChan.Close(); err != nil {
		log.Println("EmailsPublisher CloseChan: ", err)
	}
}

func (p *MessagePublisher) Publish(mensagemByte []byte, contentType string) error {
	err := rabbit.Publish(p.amqpChan, mensagemByte, config.ProductionOrderExchange, config.RoutingKey)

	if err != nil {
		return errors.Wrap(err, "ch.Publish")
	}
	return nil
}

// func (p *MessagePublisher) PublishToQueue(mensagem model.Mensagem) error {
func (p *MessagePublisher) PublishToQueue(msg string) error {
	mensagem := model.Mensagem{
		ID:   uuid.New(),
		Body: msg,
	}
	mensagemErro := model.MensagemErro{
		ID:    uuid.New(),
		Corpo: msg,
		Idade: 15,
	}
	contador++
	// if (contador % 10000) == 0 {
	if contador == 1 {
		log.Println("Mandei com erro para teste")
		mensagemBytes, err := json.Marshal(mensagemErro)
		if err != nil {
			return errors.Wrap(err, "json.Marshal")
		}
		return rabbit.Publish(p.amqpChan, mensagemBytes, config.ProductionOrderExchange, config.RoutingKey)
	} else {
		mensagemBytes, err := json.Marshal(mensagem)
		if err != nil {
			return errors.Wrap(err, "json.Marshal")
		}
		return rabbit.Publish(p.amqpChan, mensagemBytes, config.ProductionOrderExchange, config.RoutingKey)
	}
}
