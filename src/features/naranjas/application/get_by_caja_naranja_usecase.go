package application

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type GetByCajaNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewGetByDescriptionNaranjaUseCase(naranjaRepository domain.INaranja) *GetByCajaNaranjaUseCase { 
	return &GetByCajaNaranjaUseCase{naranjaRepository: naranjaRepository} 
}

func (g *GetByCajaNaranjaUseCase) Execute(cajaId int) ([]domain.Naranja, error) {
	naranjas, err := g.naranjaRepository.GetByCaja(cajaId)
	if err != nil {
		return nil, err
	}
	return naranjas, nil
}
