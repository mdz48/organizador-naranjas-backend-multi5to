package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetTop3CajasByLoteUseCase struct {
	cajaRepository domain.ICaja
}

func NewGetTop3CajasByLoteUseCase(cajaRepository domain.ICaja) *GetTop3CajasByLoteUseCase {
	return &GetTop3CajasByLoteUseCase{cajaRepository: cajaRepository}
}

func (g *GetTop3CajasByLoteUseCase) Execute(loteId int) ([]domain.Caja, error) {
	cajas, err := g.cajaRepository.GetTop3ByLote(loteId)
	if err != nil {
		return nil, err
	}
	return cajas, nil
}
