package application

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/ports"
)

type GetEsp32ByPropietarioUseCase struct {
	db ports.IEsp32
}

func NewGetEsp32ByPropietarioUseCase(db ports.IEsp32) *GetEsp32ByPropietarioUseCase {
	return &GetEsp32ByPropietarioUseCase{
		db: db,
	}
}

func (uc *GetEsp32ByPropietarioUseCase) Run(id int) (*entities.Esp32, error) {
	return uc.db.GetByPropietario(id)
}
