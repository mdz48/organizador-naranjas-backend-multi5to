package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type GetLotesByUserUseCase struct {
	loteRepository domain.ILote
}

func NewGetLotesByUserUseCase(loteRepository domain.ILote) *GetLotesByUserUseCase {
	return &GetLotesByUserUseCase{
		loteRepository: loteRepository,
	}
}

func (g *GetLotesByUserUseCase) Execute(userId int) ([]domain.Lote, error) {
	return g.loteRepository.GetByUserId(userId)
}
