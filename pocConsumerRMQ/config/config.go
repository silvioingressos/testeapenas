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
	Host     = "localhost"
	Port     = "5672"
	User     = "guest"
	Password = "guest"
	//	Exchange            = "poc-panda-exchange" // devem ser iguais
	//	DeadExchange        = "poc-panda-dead-exchange"
	//	HospitalExchange    = "poc-panda-hospital-exchange"
	//	DeadRoutingKey      = "poc-panda-dead-routing-key"
	//	Queue               = "poc-panda-queue"
	//	QueueRetry          = "poc-panda-queue-retry"
	//	Queue_hospital      = "poc-panda-queue-hospital"
	//	RoutingKey          = "poc-panda-routing-key"
	//	RoutingKey_hospital = "poc-panda-queue-hospital"
	//	RoutingKeyRetry     = "poc-panda-queue-retry"
	//	ConsumerTag         = "poc-panda-consumer"
	//	QtdRetry            = 3
	//	TimeBetweenRetry    = 10000 // 10seg

	WorkerPoolSize      = 10
	NumberOfConnections = 15

// Melhor configuracao
//	WorkerPoolSize      = 10
//	NumberOfConnections = 15

)

const (
	ProductionOrderExchange = "poc-panda-exchange"
	DeadLetterExchange      = "poc-panda-dead-exchange"
	CreateQueue             = "poc-panda-queue"
	DeadLetterQueue         = "poc-panda-queue.dlx"
	RoutingKey              = "poc-panda-routingkey"

	HospitalExchange = "poc-panda-exchange-hospital"
	HospitalQueue    = "poc-panda-queue-hospital"

	ConsumerTag = "poc-panda-consumer"
)
