package application

import "organizador-naranjas-backend-multi5to/src/features/naranjas/domain"

type DeleteNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewDeleteNaranjaOrderUseCase(caja domain.INaranja) *DeleteNaranjaUseCase {
	return &DeleteNaranjaUseCase{naranjaRepository: caja}
}

func (d *DeleteNaranjaUseCase) Execute(id int) error {
	err := d.naranjaRepository.Delete(id)
	if err != nil {
		return nil
	}
	return err
}
