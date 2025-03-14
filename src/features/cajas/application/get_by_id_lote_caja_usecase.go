package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetByIdLoteCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewGetByIdLoteCajaUseCase(cajaRepository domain.ICaja) *GetByIdLoteCajaUseCase { 
	return &GetByIdLoteCajaUseCase{cajaRepository: cajaRepository} 
}

func (g *GetByIdLoteCajaUseCase) Execute(id int) ([]domain.Caja, error) {
	cajas, err := g.cajaRepository.GetByLote(id)
	if err != nil {
		return nil, err
	}
	return cajas, nil
}
