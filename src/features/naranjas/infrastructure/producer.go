package infrastructure

import (
	"encoding/json"
	"log"
	"organizador-naranjas-backend-multi5to/src/core"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type Producer struct {
	rabbitMQ *core.RabbitMQ
}

func NewProducer(rabbitMQ *core.RabbitMQ) *Producer {
	return &Producer{
		rabbitMQ: rabbitMQ,
	}
}

func (p *Producer) PublishNaranja(naranja domain.Naranja) error {
    jsonData, err := json.Marshal(naranja)
    if err != nil {
        log.Printf("Error al convertir naranja a JSON: %v", err)
        return err
    }

    err = p.rabbitMQ.PublishMessage("api2.oranges", jsonData)
    if err != nil {
        log.Printf("Error al publicar en RabbitMQ: %v", err)
        return err
    }

    log.Printf("Naranja ID %d publicada exitosamente", naranja.ID)
    return nil
}