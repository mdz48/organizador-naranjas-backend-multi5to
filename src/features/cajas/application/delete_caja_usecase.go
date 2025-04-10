package application

import "organizador-naranjas-backend-multi5to/src/features/cajas/domain"

type DeleteCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewDeleteCajaUseCase(caja domain.ICaja) *DeleteCajaUseCase {
	return &DeleteCajaUseCase{cajaRepository: caja}
}

func (d *DeleteCajaUseCase) Execute(id int) error {
	err := d.cajaRepository.Delete(id)
	return err
}
