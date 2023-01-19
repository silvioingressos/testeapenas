package config

const (
	ExchangeKind       = "direct"
	ExchangeDurable    = true
	ExchangeAutoDelete = false
	ExchangeInternal   = false
	ExchangeNoWait     = false

	QueueDurable    = true
	QueueAutoDelete = false
	QueueExclusive  = false
	QueueNoWait     = false

	PublishMandatory = false
	PublishImmediate = false

	PrefetchCount  = 1
	PrefetchSize   = 0
	PrefetchGlobal = false

	ConsumeAutoAck   = false
	ConsumeExclusive = false
	ConsumeNoLocal   = false
	ConsumeNoWait    = false
)

const (
	Host           = "localhost"
	Port           = "5672"
	User           = "guest"
	Password       = "guest"
	WorkerPoolSize = 23
)

const (
	ProductionOrderExchange = "poc-panda-exchange"
	DeadLetterExchange      = "poc-panda-dead-exchange"
	CreateQueue             = "poc-panda-queue"
	DeadLetterQueue         = "poc-panda-queue.dlx"
	RoutingKey              = "poc-panda-routingkey"

	HospitalExchange = "poc-panda-exchange-hospital"
	HospitalQueue    = "poc-panda-queue-hospital"
)
