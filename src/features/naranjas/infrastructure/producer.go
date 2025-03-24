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
	// Verificar si rabbitMQ es nil antes de usarlo
	if p.rabbitMQ == nil {
		log.Printf("Advertencia: RabbitMQ no está configurado, saltando publicación")
		return nil // No consideramos esto un error fatal
	}

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

	log.Printf("Naranja ID %d publicada exitosamente en RabbitMQ", naranja.ID)
	return nil
}
