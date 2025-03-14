package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetAllUseCase struct {
	cajaRepository domain.ICaja
}

func NewGetAllUseCase(cajaRepository domain.ICaja) *GetAllUseCase { return &GetAllUseCase{cajaRepository: cajaRepository} }

func (g *GetAllUseCase) Execute() ([]domain.Caja, error) {
	cajas, err := g.cajaRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return cajas, nil
}
