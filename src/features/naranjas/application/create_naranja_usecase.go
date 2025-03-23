package application

import (
    "log"
    "organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type CreateNaranjaUseCase struct {
    naranjaRepository domain.INaranja
    producer         domain.IProducer
}

func NewCreateNaranjaUseCase(
    naranjaRepository domain.INaranja,
    producer domain.IProducer,
) *CreateNaranjaUseCase {
    return &CreateNaranjaUseCase{
        naranjaRepository: naranjaRepository,
        producer:         producer,
    }
}

func (c *CreateNaranjaUseCase) Execute(naranja domain.Naranja) (domain.Naranja, error) {
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
