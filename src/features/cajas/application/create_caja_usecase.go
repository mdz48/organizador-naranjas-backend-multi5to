package application

import "organizador-naranjas-backend-multi5to/src/features/cajas/domain"

type CreateCajaUseCase struct {
	cajaRepository domain.ICaja
}

func NewCreateCajaUseCase(cajaRepository domain.ICaja) *CreateCajaUseCase {
	return &CreateCajaUseCase{cajaRepository: cajaRepository}
}

func (c *CreateCajaUseCase) Execute(caja domain.Caja) (domain.Caja, error) {
	cajaCreada, err := c.cajaRepository.Create(caja)
	if err != nil {
		return domain.Caja{}, err
	}
	return cajaCreada, nil
}
