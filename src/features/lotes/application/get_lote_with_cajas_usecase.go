package application

import (
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type GetLoteWithCajasUseCase struct {
	loteRepository domain.ILote
	cajaRepository cajaDomain.ICaja
}

func NewGetLoteWithCajasUseCase(loteRepository domain.ILote, cajaRepository cajaDomain.ICaja) *GetLoteWithCajasUseCase {
	return &GetLoteWithCajasUseCase{
		loteRepository: loteRepository,
		cajaRepository: cajaRepository,
	}
}

func (g *GetLoteWithCajasUseCase) Execute(loteId int) (domain.LoteWithCajasResponse, error) {
	// 1. Obtener el lote
	lote, err := g.loteRepository.GetById(loteId)
	if err != nil {
		return domain.LoteWithCajasResponse{}, err
	}

	// 2. Obtener las cajas asociadas al lote
	cajas, err := g.cajaRepository.GetTop3ByLote(loteId)
	if err != nil {
		return domain.LoteWithCajasResponse{}, err
	}

	// Verificar si cajas es nil y usar array vacío en ese caso
	if cajas == nil {
		cajas = []cajaDomain.Caja{}
	}

	// 3. Crear y devolver la respuesta
	response := domain.LoteWithCajasResponse{
		Lote:  lote,
		Cajas: cajas,
	}

	return response, nil
}
