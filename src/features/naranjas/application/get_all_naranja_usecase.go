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

func (u *GetAllNaranjaUseCase) Execute() ([]domain.Naranja, error) {
	return u.naranjaRepository.GetAll()
}