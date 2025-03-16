package application

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type ListLotesUseCase struct {
	loteRepository domain.ILote
}

func NewListLotesUseCase(loteRepository domain.ILote) *ListLotesUseCase {
	return &ListLotesUseCase{
		loteRepository: loteRepository,
	}
}

func (uc *ListLotesUseCase) Run() ([]domain.Lote, error) {
	lotes, err := uc.loteRepository.GetAll();

	if  err != nil {
		return nil, err; 
	}
	return lotes, nil; 
}