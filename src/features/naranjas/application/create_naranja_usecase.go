package application

import (
	"errors"
	"log"
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
	"time"
)

type CreateNaranjaUseCase struct {
	naranjaRepository domain.INaranja
	producer          domain.IProducer
	findCajaUseCase   *application.FindActiveCajaByEsp32UseCase
}

func NewCreateNaranjaUseCase(
	naranjaRepository domain.INaranja,
	producer domain.IProducer,
	findCajaUseCase *application.FindActiveCajaByEsp32UseCase,
) *CreateNaranjaUseCase {
	return &CreateNaranjaUseCase{
		naranjaRepository: naranjaRepository,
		producer:          producer,
		findCajaUseCase:   findCajaUseCase,
	}
}

func (c *CreateNaranjaUseCase) Execute(naranja domain.Naranja) (domain.Naranja, error) {
	if naranja.Hora.IsZero() {
		naranja.Hora = time.Now()
	}

	if naranja.CajaFK == 0 && naranja.Esp32FK != "" {
		// Validar que se proporcione un color
		if naranja.Color == "" {
			return domain.Naranja{}, errors.New("el color de la naranja es requerido para asignarla a una caja")
		}

		// Buscar una caja activa para este ESP32 que coincida con el color de la naranja
		caja, err := c.findCajaUseCase.Execute(naranja.Esp32FK, naranja.Color)
		if err != nil {
			return domain.Naranja{}, errors.New("no se encontró una caja activa para este ESP32 y color: " + err.Error())
		}

		naranja.CajaFK = caja.ID
		log.Printf("Naranja de color %s automáticamente asignada a la caja #%d: %s",
			naranja.Color, caja.ID, caja.Descripcion)
	} else if naranja.CajaFK == 0 {
		return domain.Naranja{}, errors.New("se requiere especificar esp32_fk o caja_fk")
	}

	naranjaCreada, err := c.naranjaRepository.Create(naranja)
	if err != nil {
		return domain.Naranja{}, err
	}

	// Publicar evento en RabbitMQ
	if c.producer != nil {
		if err := c.producer.PublishNaranja(naranjaCreada); err != nil {
			log.Printf("Error al publicar en RabbitMQ: %v", err)
		}
	}

	return naranjaCreada, nil
}

// ALERTA!!! EN LA BD TRIGGER A CAJA
