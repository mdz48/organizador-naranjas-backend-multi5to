package application

import "organizador-naranjas-backend-multi5to/src/features/naranjas/domain"

type CreateNaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewCreateCajaUseCase(naranjaRepository domain.INaranja) *CreateNaranjaUseCase {
	return &CreateNaranjaUseCase{naranjaRepository: naranjaRepository}
}

func (c *CreateNaranjaUseCase) Execute(naranja domain.Naranja) (domain.Naranja, error) {
	naranjaCreada, err := c.naranjaRepository.Create(naranja)
	if err != nil {
		return domain.Naranja{}, err
	}
	return naranjaCreada, nil
}
