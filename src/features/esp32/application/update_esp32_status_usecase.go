package application

import (
	"fmt"
	"log"
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
	loteUseCase "organizador-naranjas-backend-multi5to/src/features/lotes/application"
)

type UpdateEsp32StatusUseCase struct {
	db                      ports.IEsp32
	cajaRepository          cajaDomain.ICaja
	updateLoteStatusUseCase *loteUseCase.UpdateLoteStatusUseCase
	producer                ports.IEsp32Producer
}

func NewUpdateEsp32StatusUseCase(
	db ports.IEsp32,
	cajaRepository cajaDomain.ICaja,
	updateLoteStatusUseCase *loteUseCase.UpdateLoteStatusUseCase,
	producer ports.IEsp32Producer,
) *UpdateEsp32StatusUseCase {
	return &UpdateEsp32StatusUseCase{
		db:                      db,
		cajaRepository:          cajaRepository,
		updateLoteStatusUseCase: updateLoteStatusUseCase,
		producer:                producer,
	}
}

func (uc *UpdateEsp32StatusUseCase) Run(id string, status string) error {
	// Validar que el status sea válido
	if status != "esperando" && status != "activo" && status != "desactivado" {
		return fmt.Errorf("status inválido: %s (debe ser 'esperando', 'activo' o 'desactivado')", status)
	}

	// Si el estado es "esperando" o "desactivado", terminar los lotes asociados
	if status == "esperando" || status == "desactivado" {
		loteIds, err := uc.cajaRepository.GetLotesByEsp32(id, "CARGANDO")
		if err != nil {
			log.Printf("Error al buscar lotes por ESP32 %s: %v", id, err)
			// No fallamos la actualización del ESP32, solo registramos el error
		}

		// Actualizar cada lote encontrado a estado "terminado"
		for _, loteId := range loteIds {
			_, err := uc.updateLoteStatusUseCase.Run(loteId, "terminado")
			if err != nil {
				log.Printf("Error al actualizar lote %d: %v", loteId, err)
				// Continuamos con otros lotes
			} else {
				log.Printf("Lote %d marcado como terminado debido al cambio de estado del ESP32 %s", loteId, id)
			}
		}
	}

	// Actualizar el estado del ESP32 en la base de datos
	err := uc.db.UpdateStatus(id, status)
	if err != nil {
		return err
	}

	// Publicar el mensaje de cambio de estado
	if uc.producer != nil {
		if err := uc.producer.PublishStatusChange(id, status); err != nil {
			log.Printf("Error al publicar mensaje de cambio de estado para ESP32 %s: %v", id, err)
			// No fallamos la operación si la publicación del mensaje falla
		}
	}

	return nil
}
