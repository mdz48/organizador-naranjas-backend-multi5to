package application

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type GetByIdNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewGetByIdUseCase(naranjaRepository domain.INaranja) *GetByIdNaranjaUseCase { 
	return &GetByIdNaranjaUseCase{naranjaRepository: naranjaRepository} 
}

func (g *GetByIdNaranjaUseCase) Execute(id int)  {
	/*
	naranja, err := g.naranjaRepository.GetById(id)
	if err != nil {
		return domain.Naranja{}, err
	}
	return naranja, nil
	*/
	}
