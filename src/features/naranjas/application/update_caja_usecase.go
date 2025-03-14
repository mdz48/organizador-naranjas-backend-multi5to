package application

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type UpdateNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewUpdateNaranjaUseCase(naranjaRepository domain.INaranja) *UpdateNaranjaUseCase {
	return &UpdateNaranjaUseCase{naranjaRepository: naranjaRepository}
}

func (u *UpdateNaranjaUseCase) Execute(naranja domain.Naranja) {
	/*
	caja, err := u.naranjaRepository.Update(naranja)
	if err != nil {
		return domain.Naranja{}, err
	}
	return caja, nil
	*/
	}
