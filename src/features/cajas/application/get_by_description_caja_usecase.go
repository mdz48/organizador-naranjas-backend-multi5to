package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetByDescriptionCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewGetByDescriptionCajaUseCase(cajaRepository domain.ICaja) *GetByDescriptionCajaUseCase { 
	return &GetByDescriptionCajaUseCase{cajaRepository: cajaRepository} 
}

func (g *GetByDescriptionCajaUseCase) Execute(descripcion string) (domain.Caja, error) {
	caja, err := g.cajaRepository.GetByDescripcion(descripcion)
	if err != nil {
		return domain.Caja{}, err
	}
	return caja, nil
}
