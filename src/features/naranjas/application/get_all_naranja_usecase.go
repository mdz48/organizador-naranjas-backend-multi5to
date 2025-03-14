package application

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type GetAllNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewGetAllUseCase(naranjaRepository domain.INaranja) *GetAllNaranjaUseCase { 
	return &GetAllNaranjaUseCase{naranjaRepository: naranjaRepository} 
}

func (g *GetAllNaranjaUseCase) Execute() ([]domain.Naranja, error) {
	cajas, err := g.naranjaRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return cajas, nil
}
