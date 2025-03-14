package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type CreateLoteUseCase struct {
	loteRepository domain.ILote
}

func NewCreateLoteUseCase(loteRepository domain.ILote) *CreateLoteUseCase {
	return &CreateLoteUseCase{loteRepository: loteRepository}
}

func (c *CreateLoteUseCase) Execute(lote domain.Lote) (domain.Lote, error) {
	return c.loteRepository.Create(lote)
}
