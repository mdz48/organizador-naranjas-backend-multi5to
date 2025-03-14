package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetByIdCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewGetByIdUseCase(cajaRepository domain.ICaja) *GetByIdCajaUseCase { 
	return &GetByIdCajaUseCase{cajaRepository: cajaRepository} 
}

func (g *GetByIdCajaUseCase) Execute(id int) (domain.Caja, error) {
	caja, err := g.cajaRepository.GetById(id)
	if err != nil {
		return domain.Caja{}, err
	}
	return caja, nil
}
