package application

import (
	"log"
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type GetAllLotesWithCajasUseCase struct {
    loteRepository domain.ILote
    cajaRepository cajaDomain.ICaja
}

func NewGetAllLotesWithCajasUseCase(loteRepository domain.ILote, cajaRepository cajaDomain.ICaja) *GetAllLotesWithCajasUseCase {
    return &GetAllLotesWithCajasUseCase{
        loteRepository: loteRepository,
        cajaRepository: cajaRepository,
    }
}

func (g *GetAllLotesWithCajasUseCase) Execute() ([]domain.LoteWithCajasResponse, error) {
    // 1. Obtener todos los lotes
    lotes, err := g.loteRepository.GetAll()
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
            continue
        }

        loteWithCajas := domain.LoteWithCajasResponse{
            Lote:  lote,
            Cajas: cajas,
        }

        response = append(response, loteWithCajas)
    }

    return response, nil
}