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


▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
debug
{"LEVEL":"debug","TIME":"Jan 03 19:28:32.746220900","CALLER":"logs/main.go:10","MESSAGE":"Debug -> logger construction succeeded"}
{"LEVEL":"info","TIME":"Jan 03 19:28:32.746746200","CALLER":"logs/main.go:11","MESSAGE":"Info -> logger construction succeeded"}
{"LEVEL":"warn","TIME":"Jan 03 19:28:32.746746200","CALLER":"logs/main.go:12","MESSAGE":"Warn -> logger construction succeeded"}
{"LEVEL":"error","TIME":"Jan 03 19:28:32.746812400","CALLER":"logs/main.go:13","MESSAGE":"Error -> logger construction succeeded"}
{"LEVEL":"dpanic","TIME":"Jan 03 19:28:32.746812400","CALLER":"logs/main.go:14","MESSAGE":"dpanic -> logger construction succeeded"}
{"LEVEL":"fatal","TIME":"Jan 03 19:28:32.746812400","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}

info
{"LEVEL":"info","TIME":"Jan 03 19:28:57.672162700","CALLER":"logs/main.go:11","MESSAGE":"Info -> logger construction succeeded"}
{"LEVEL":"warn","TIME":"Jan 03 19:28:57.672162700","CALLER":"logs/main.go:12","MESSAGE":"Warn -> logger construction succeeded"}
{"LEVEL":"error","TIME":"Jan 03 19:28:57.672162700","CALLER":"logs/main.go:13","MESSAGE":"Error -> logger construction succeeded"}
{"LEVEL":"dpanic","TIME":"Jan 03 19:28:57.672162700","CALLER":"logs/main.go:14","MESSAGE":"dpanic -> logger construction succeeded"}
{"LEVEL":"fatal","TIME":"Jan 03 19:28:57.672162700","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}

warn
{"LEVEL":"warn","TIME":"Jan 03 19:29:23.462454400","CALLER":"logs/main.go:12","MESSAGE":"Warn -> logger construction succeeded"}
{"LEVEL":"error","TIME":"Jan 03 19:29:23.462454400","CALLER":"logs/main.go:13","MESSAGE":"Error -> logger construction succeeded"}
{"LEVEL":"dpanic","TIME":"Jan 03 19:29:23.462454400","CALLER":"logs/main.go:14","MESSAGE":"dpanic -> logger construction succeeded"}
{"LEVEL":"fatal","TIME":"Jan 03 19:29:23.462454400","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}

error
{"LEVEL":"error","TIME":"Jan 03 19:29:50.801097900","CALLER":"logs/main.go:13","MESSAGE":"Error -> logger construction succeeded"}
{"LEVEL":"dpanic","TIME":"Jan 03 19:29:50.801728600","CALLER":"logs/main.go:14","MESSAGE":"dpanic -> logger construction succeeded"}
{"LEVEL":"fatal","TIME":"Jan 03 19:29:50.801832800","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}

dpanic
{"LEVEL":"dpanic","TIME":"Jan 03 19:30:09.493287600","CALLER":"logs/main.go:14","MESSAGE":"dpanic -> logger construction succeeded"}
{"LEVEL":"fatal","TIME":"Jan 03 19:30:09.493287600","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}

panic


fatal
{"LEVEL":"fatal","TIME":"Jan 03 19:30:28.571481100","CALLER":"logs/main.go:16","MESSAGE":"Fatal -> logger construction succeeded"}
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
