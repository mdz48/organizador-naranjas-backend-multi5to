// Nuevo archivo: src/features/lotes/application/create_lote_with_cajas_usecase.go
package application

import (
    "time"
    cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
    cajaRepo "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
    "organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type CreateLoteWithCajasUseCase struct {
    loteRepository domain.ILote
    cajaRepository cajaRepo.ICaja
}

func NewCreateLoteWithCajasUseCase(loteRepository domain.ILote, cajaRepository cajaRepo.ICaja) *CreateLoteWithCajasUseCase {
    return &CreateLoteWithCajasUseCase{
        loteRepository: loteRepository,
        cajaRepository: cajaRepository,
    }
}

func (c *CreateLoteWithCajasUseCase) Execute(lote domain.Lote, esp32FK string) (domain.Lote, []cajaDomain.Caja, error) {

    if lote.Fecha == "" {
        lote.Fecha = time.Now().Format("2006-01-02")
    }

    if lote.Estado == "" {
        lote.Estado = "cargando"
    }

    createdLote, err := c.loteRepository.Create(lote)
    if err != nil {
        return domain.Lote{}, nil, err
    }

    // Crear las 3 cajas asociadas a este lote
    cajas := make([]cajaDomain.Caja, 0, 3)
    cajaDescripciones := []string{"naranja", "verde", "maduracion"}
    
    ahora := time.Now()
    for _, descripcion := range cajaDescripciones {
        caja := cajaDomain.Caja{
            Descripcion: descripcion,
            PesoTotal:   0,
            Precio:      0,
            HoraInicio:  ahora,
            LoteFK:      createdLote.ID,
            EncargadoFK: lote.UserID, 
            Cantidad:    0,
            Estado:      "CARGANDO",
			Esp32FK:     esp32FK,
        }
        
        createdCaja, err := c.cajaRepository.Create(caja)
        if err != nil {
            return domain.Lote{}, nil, err
        }
        
        cajas = append(cajas, createdCaja)
    }

    return createdLote, cajas, nil
}