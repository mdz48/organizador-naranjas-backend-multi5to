// Nuevo archivo: src/features/lotes/application/create_lote_with_cajas_usecase.go
package application

import (
	"fmt"
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	cajaRepo "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
	messagePort "organizador-naranjas-backend-multi5to/src/features/lotes/domain/ports"
	"time"
)

type CreateLoteWithCajasUseCase struct {
	loteRepository  domain.ILote
	cajaRepository  cajaRepo.ICaja
	esp32Repository ports.IEsp32
	messageProducer messagePort.IMessageProducer
}

func NewCreateLoteWithCajasUseCase(
	loteRepository domain.ILote,
	cajaRepository cajaRepo.ICaja,
	esp32Repository ports.IEsp32,
	messageProducer messagePort.IMessageProducer,
) *CreateLoteWithCajasUseCase {
	return &CreateLoteWithCajasUseCase{
		loteRepository:  loteRepository,
		cajaRepository:  cajaRepository,
		esp32Repository: esp32Repository,
		messageProducer: messageProducer,
	}
}

func (c *CreateLoteWithCajasUseCase) Execute(lote domain.Lote, esp32FK string) (domain.Lote, []cajaDomain.Caja, error) {
	// Si se proporcionó un ID de ESP32, verificar que exista y que esté en estado "esperando"
	if esp32FK != "" {
		esp32, err := c.esp32Repository.GetById(esp32FK)
		if err != nil {
			return domain.Lote{}, nil, fmt.Errorf("error al obtener ESP32 con ID %s: %v", esp32FK, err)
		}

		// Verificar el estado del ESP32
		if esp32.Status != "esperando" {
			return domain.Lote{}, nil, fmt.Errorf("el ESP32 con ID %s tiene estado '%s'. Solo se pueden crear lotes con ESP32 en estado 'esperando'", esp32FK, esp32.Status)
		}
	}

	if lote.Fecha == "" {
		lote.Fecha = time.Now().Format("2006-01-02")
	}

	if lote.Estado == "" {
		lote.Estado = "cargando"
	}

	createdLote, err := c.loteRepository.Create(lote)
	if err != nil {
		return domain.Lote{}, nil, err
	}

	// Crear las 3 cajas asociadas a este lote
	cajas := make([]cajaDomain.Caja, 0, 3)
	cajaDescripciones := []string{"naranja", "verde", "maduracion"}

	ahora := time.Now()
	for _, descripcion := range cajaDescripciones {
		caja := cajaDomain.Caja{
			Descripcion: descripcion,
			PesoTotal:   0,
			Precio:      0,
			HoraInicio:  ahora,
			HoraFin:     nil,
			LoteFK:      createdLote.ID,
			EncargadoFK: lote.UserID,
			Cantidad:    0,
			Estado:      "CARGANDO",
			Esp32FK:     esp32FK,
		}

		createdCaja, err := c.cajaRepository.Create(caja)
		if err != nil {
			return domain.Lote{}, nil, err
		}

		cajas = append(cajas, createdCaja)
	}

	// Si se asignó un ESP32, actualizar su estado a "activo"
	if esp32FK != "" {
		err = c.esp32Repository.UpdateStatus(esp32FK, "activo")
		if err != nil {
			// Loguear el error pero no fallar la creación
			fmt.Printf("Error al actualizar el estado del ESP32 %s: %v\n", esp32FK, err)
		}

		// Enviar mensaje a RabbitMQ para notificar a la ESP32
		message := messagePort.Message{
			Esp32FK: esp32FK,
			Content: "esperando",
		}

		err = c.messageProducer.SendMessage(message)
		if err != nil {
			// Loguear el error pero no fallar la creación
			fmt.Printf("Error al enviar mensaje a la ESP32 %s: %v\n", esp32FK, err)
		}
	}

	return createdLote, cajas, nil
}
