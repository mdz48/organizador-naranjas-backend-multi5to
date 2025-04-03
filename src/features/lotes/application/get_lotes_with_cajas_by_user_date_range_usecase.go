package application

import (
	"log"
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type GetLotesWithCajasByUserDateRangeUseCase struct {
	loteRepository domain.ILote
	cajaRepository cajaDomain.ICaja
}

func NewGetLotesWithCajasByUserDateRangeUseCase(loteRepository domain.ILote, cajaRepository cajaDomain.ICaja) *GetLotesWithCajasByUserDateRangeUseCase {
	return &GetLotesWithCajasByUserDateRangeUseCase{
		loteRepository: loteRepository,
		cajaRepository: cajaRepository,
	}
}

func (g *GetLotesWithCajasByUserDateRangeUseCase) Execute(userId int, startDate, endDate string) ([]domain.LoteWithCajasResponse, error) {
	// 1. Obtener los lotes del usuario en el rango de fechas especificado
	lotes, err := g.loteRepository.GetByUserIdAndDateRange(userId, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// 2. Crear slice para almacenar las respuestas
	response := make([]domain.LoteWithCajasResponse, 0, len(lotes))

	// 3. Para cada lote, buscar sus cajas asociadas
	for _, lote := range lotes {
		cajas, err := g.cajaRepository.GetByLote(lote.ID)
		if err != nil {
			// Loguear el error pero continuar con otros lotes
			log.Printf("Error al obtener cajas para lote %d: %v", lote.ID, err)
			// Usar array vacío en lugar de null/nil
			cajas = []cajaDomain.Caja{}
		}

		// Si no hay cajas, usar un array vacío
		if cajas == nil {
			cajas = []cajaDomain.Caja{}
		}

		loteWithCajas := domain.LoteWithCajasResponse{
			Lote:  lote,
			Cajas: cajas,
		}

		response = append(response, loteWithCajas)
	}

	return response, nil
}
