main

appLoggers := make([]*logger.ApiLogger, sizePool)

s := server.NewServer(amqpConnConsumers[i], amqpConnProducers[i], appLoggers[i], dynamoDbs[i])

------
server

type Server struct {
	dynamoDB         *mongo.Client
	amqpConnConsumer *amqp.Connection
	amqpConnProducer *amqp.Connection
	log              logger.Logger
}

s.log.Error("StartConsumer: %v", err)

------------------------
repository

type Repository struct {
	dynamoDB *mongo.Client
	log      logger.Logger
	cronus   *util.Cronometro
}
func NewRepository(db *mongo.Client, logger logger.Logger, cronus *util.Cronometro) *Repository {
	return &Repository{
		dynamoDB: db,
		log:      logger,
		cronus:   cronus,
	}
}

func (r *Repository) GetFundoExternoEntidadeRV(cge int64) ([]model.EntidadeFundos, error) {
	inicio := time.Now()
	defer r.cronus.Chronometer("Repository -> GetFundoExternoEntidadeRV", &inicio)