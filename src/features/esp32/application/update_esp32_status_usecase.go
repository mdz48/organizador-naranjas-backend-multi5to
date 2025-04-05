package application

import (
	"fmt"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
)

type UpdateEsp32StatusUseCase struct {
	db ports.IEsp32
}

func NewUpdateEsp32StatusUseCase(db ports.IEsp32) *UpdateEsp32StatusUseCase {
	return &UpdateEsp32StatusUseCase{
		db: db,
	}
}

func (uc *UpdateEsp32StatusUseCase) Run(id string, status string) error {
	// Validar que el status sea válido
	if status != "esperando" && status != "activo" && status != "desactivado" {
		return fmt.Errorf("status inválido: %s (debe ser 'esperando', 'activo' o 'desactivado')", status)
	}

	return uc.db.UpdateStatus(id, status)
}
