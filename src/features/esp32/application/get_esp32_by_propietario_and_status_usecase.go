package application

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
)

type GetEsp32ByPropietarioAndStatusUseCase struct {
	db ports.IEsp32
}

func NewGetEsp32ByPropietarioAndStatusUseCase(db ports.IEsp32) *GetEsp32ByPropietarioAndStatusUseCase {
	return &GetEsp32ByPropietarioAndStatusUseCase{
		db: db,
	}
}

func (uc *GetEsp32ByPropietarioAndStatusUseCase) Run(id int, status string) ([]entities.Esp32, error) {
	return uc.db.GetByPropietarioAndStatus(id, status)
}
