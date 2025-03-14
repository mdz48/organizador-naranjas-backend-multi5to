package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type UpdateCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewUpdateCajaUseCase(caja domain.ICaja) *UpdateCajaUseCase {
	return &UpdateCajaUseCase{cajaRepository: caja}
}

func (u *UpdateCajaUseCase) Execute(caja domain.Caja) (domain.Caja, error) {
	caja, err := u.cajaRepository.Update(caja)
	if err != nil {
		return domain.Caja{}, err
	}
	return caja, nil
}
