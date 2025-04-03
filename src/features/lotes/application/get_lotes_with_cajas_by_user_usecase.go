package application

import (
    "log"
    cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
    "organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type GetLotesWithCajasByUserUseCase struct {
    loteRepository domain.ILote
    cajaRepository cajaDomain.ICaja
}

func NewGetLotesWithCajasByUserUseCase(loteRepository domain.ILote, cajaRepository cajaDomain.ICaja) *GetLotesWithCajasByUserUseCase {
    return &GetLotesWithCajasByUserUseCase{
        loteRepository: loteRepository,
        cajaRepository: cajaRepository,
    }
}

func (g *GetLotesWithCajasByUserUseCase) Execute(userId int) ([]domain.LoteWithCajasResponse, error) {
    // 1. Obtener todos los lotes del usuario
    lotes, err := g.loteRepository.GetByUserId(userId)
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
