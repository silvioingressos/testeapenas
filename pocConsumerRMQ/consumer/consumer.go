package consumer

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"panda-poc-consumer-fila/config"
	"panda-poc-consumer-fila/model"
	rabbitlocal "panda-poc-consumer-fila/rabbit"

	// rabbitlocal "panda-poc-consumer-fila/rabbit"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Consumer struct {
	amqpConn *amqp.Connection // Channel Consumer
	chP      *amqp.Channel    // Channel Producer
	chH      *amqp.Channel    // Channel Producer manda para hospital
}

func NewImagesConsumer(amqpConn *amqp.Connection) *Consumer {
	return &Consumer{amqpConn: amqpConn}
}

func (c *Consumer) CreateChannel(exchangeName, queueName, bindingKey, consumerTag string) (*amqp.Channel, *amqp.Channel, *amqp.Channel, error) {
	ch, err := c.amqpConn.Channel()
	chP, err := c.amqpConn.Channel()
	chH, err := c.amqpConn.Channel()
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Error amqpConn.Channel")
	}

	err = ch.Qos(
		config.PrefetchCount,  // prefetch count
		config.PrefetchSize,   // prefetch size
		config.PrefetchGlobal, // global
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Error  ch.Qos")
	}

	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Error Channel Producer")
	}

	err = chP.Qos(
		config.PrefetchCount,  // prefetch count
		config.PrefetchSize,   // prefetch size
		config.PrefetchGlobal, // global
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Error  ch.Qos")
	}

	err = chH.Qos(
		config.PrefetchCount,  // prefetch count
		config.PrefetchSize,   // prefetch size
		config.PrefetchGlobal, // global
	)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "Error  ch.Qos")
	}

	/*
		err = chP.ExchangeDeclare(
			config.DeadExchange,
			config.ExchangeKind,
			config.ExchangeDurable,
			config.ExchangeAutoDelete,
			config.ExchangeInternal,
			config.ExchangeNoWait,
			nil,
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
		}

		queueP, err := chP.QueueDeclare(
			config.QueueRetry,
			config.QueueDurable,
			config.QueueAutoDelete,
			config.QueueExclusive,
			config.QueueNoWait,
			amqp.Table{
				"x-message-ttl":             16000,
				"x-dead-letter-exchange":    config.Exchange,
				"x-dead-letter-routing-key": config.RoutingKey,
			},
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.QueueDeclare")
		}

		err = chP.QueueBind(
			queueP.Name,
			config.RoutingKey,
			config.Exchange,
			config.QueueNoWait,
			nil,
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.QueueBind")
		}

		err = chP.Qos(
			config.PrefetchCount,  // prefetch count
			config.PrefetchSize,   // prefetch size
			config.PrefetchGlobal, // global
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error  ch.Qos")
		}

		//// Channel Hospital
		chH, err := c.amqpConn.Channel()
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error Channel Producer")
		}
		err = chH.ExchangeDeclare(
			config.HospitalExchange,
			config.ExchangeKind,
			config.ExchangeDurable,
			config.ExchangeAutoDelete,
			config.ExchangeInternal,
			config.ExchangeNoWait,
			nil,
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.ExchangeDeclare")
		}

		queueH, err := chP.QueueDeclare(
			config.Queue_hospital, // <--- Atencao aqui
			config.QueueDurable,
			config.QueueAutoDelete,
			config.QueueExclusive,
			config.QueueNoWait,
			nil,
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.QueueDeclare")
		}

		err = chH.QueueBind(
			queueH.Name,
			config.RoutingKey_hospital,
			config.HospitalExchange,
			config.QueueNoWait,
			nil,
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error ch.QueueBind")
		}

		// log.Println("Queue bound to exchange, starting to consume from queue, consumerTag: %v", consumerTag)

		err = chH.Qos(
			config.PrefetchCount,  // prefetch count
			config.PrefetchSize,   // prefetch size
			config.PrefetchGlobal, // global
		)
		if err != nil {
			return nil, nil, nil, errors.Wrap(err, "Error  ch.Qos")
		}
	*/
	return ch, chP, chH, nil
}

func (c *Consumer) worker(ctx context.Context, chP *amqp.Channel, chH *amqp.Channel, messages <-chan amqp.Delivery) {

	for delivery := range messages {

		err := c.Alguma_Chamada_Externa(ctx, chP, chH, delivery)
		if err != nil {
			if err := delivery.Reject(false); err != nil {
				log.Println("Err delivery.Reject: ", err)
			}
			log.Println("Failed to process delivery: ", err)

		} else {
			err = delivery.Ack(false)
			if err != nil {
				log.Println("Failed to acknowledge delivery: ", err)
			}
		}
	}
	log.Println("Deliveries channel closed")
}

// StartConsumer Start new rabbitmq consumer
func (c *Consumer) StartConsumer(workerPoolSize int, exchange, queueName, bindingKey, consumerTag string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch, chP, chH, err := c.CreateChannel(exchange, queueName, bindingKey, consumerTag)
	if err != nil {
		return errors.Wrap(err, "CreateChannel")
	}
	defer ch.Close()
	defer chP.Close()
	defer chH.Close()

	deliveries, err := ch.Consume(
		queueName,
		consumerTag,
		config.ConsumeAutoAck,
		config.ConsumeExclusive,
		config.ConsumeNoLocal,
		config.ConsumeNoWait,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	for i := 0; i < workerPoolSize; i++ {
		go c.worker(ctx, chP, chH, deliveries)
	}

	chanErr := <-ch.NotifyClose(make(chan *amqp.Error))
	log.Println("ch.NotifyClose: ", chanErr)
	return chanErr
}

func (c *Consumer) handleFailedMsg(chP *amqp.Channel, chH *amqp.Channel, m amqp.Delivery) {

	retryCount := c.getRetry(m.Headers)
	headers := make(amqp.Table)
	if retryCount >= 3 {
		rabbitlocal.Publish(chH, m.Body, config.HospitalExchange, config.RoutingKey, nil)
		log.Println("Mensagem foi para o hospital")
	} else {
		headers["x-retry-count"] = retryCount + 1
		rabbitlocal.Publish(chP, m.Body, config.DeadLetterExchange, config.RoutingKey, headers)
	}
}

func (c *Consumer) getDelay(h amqp.Table) int32 {
	var d int32
	if h["x-delay"] != nil {
		n := h["x-delay"].(int32)
		d = int32(math.Abs(float64(n)))
	}
	return d
}

// getRetry returns the retry count from the header
func (c *Consumer) getRetry(h amqp.Table) int32 {
	var r int32
	lastCount := h["x-retry-count"]
	if lastCount != nil {
		r = lastCount.(int32)
	}
	return r
}

func (c *Consumer) Alguma_Chamada_Externa(ctx context.Context, chP *amqp.Channel, chH *amqp.Channel, m amqp.Delivery) error {
	data := model.Mensagem{}
	json.Unmarshal(m.Body, &data)
	texto, err := json.Marshal(data)
	if err != nil {
		log.Println("Temos o erro ", err)
		c.handleFailedMsg(chP, chP, m)
		return err
	}
	if data.Body == "" {
		log.Println("### ### ### Erro de parser (Proposito) ")
		c.handleFailedMsg(chP, chP, m)
	}

	log.Println(" Mensagem: ", string(texto))
	return nil
}
