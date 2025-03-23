package application

import "organizador-naranjas-backend-multi5to/src/features/naranjas/domain"

type DeleteNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewDeleteNaranjaOrderUseCase(caja domain.INaranja) *DeleteNaranjaUseCase {
	return &DeleteNaranjaUseCase{naranjaRepository: caja}
}

func (d *DeleteNaranjaUseCase) Execute(id int)  {
	d.naranjaRepository.Delete(id)
}