package util

import (
	"Panda-DockerGo-Consolidacao-Posicao-Entidade/logger"
	"log"
	"time"
)

type Cronometro struct {
	log logger.Logger
}

func NewChronometer(logger logger.Logger) *Cronometro {
	return &Cronometro{
		log: logger,
	}
}

func (c *Cronometro) Chronometer(mensagem string, inicio *time.Time) {
	fim := time.Now()
	tempo_execucao := fim.Sub(*inicio)
	if tempo_execucao == 0 {
		log.Println("Chronometer -> ", mensagem, " : ", tempo_execucao.Milliseconds(), " millesegundos")
		return
	} else {
		c.log.DPanic(mensagem, " : ", tempo_execucao.Seconds(), " segundos")
	}

}
