package application

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type GetByEsp32NaranjaUseCase struct {
	naranjaRepository domain.INaranja
}

func NewGetByEsp32NaranjaUseCase(naranjaRepository domain.INaranja) *GetByEsp32NaranjaUseCase {
	return &GetByEsp32NaranjaUseCase{naranjaRepository: naranjaRepository}
}

func (g *GetByEsp32NaranjaUseCase) Execute(esp32Id string) ([]domain.Naranja, error) {
	naranjas, err := g.naranjaRepository.GetByEsp32(esp32Id)
	if err != nil {
		return nil, err
	}
	return naranjas, nil
}
