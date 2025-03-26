package application

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
	"time"
)

type CreateLoteUseCase struct {
	loteRepository domain.ILote
}

func NewCreateLoteUseCase(loteRepository domain.ILote) *CreateLoteUseCase {
	return &CreateLoteUseCase{loteRepository: loteRepository}
}

func (c *CreateLoteUseCase) Execute(lote domain.Lote) (domain.Lote, error) {
	if lote.Fecha == "" {
		lote.Fecha = time.Now().Format("2006-01-02")
	}

	if lote.Estado == "" {
		lote.Estado = "cargando"
	}

	return c.loteRepository.Create(lote)
}
