package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type ListLoteDateUseCase struct {
	loteRepository domain.ILote
}

func NewListLoteDateUseCase(loteRepository domain.ILote) *ListLoteDateUseCase {
	return &ListLoteDateUseCase{
		loteRepository: loteRepository,
	}
}

func (uc *ListLoteDateUseCase) Run(date string) ([]domain.Lote, error) {
	lote, err := uc.loteRepository.GetByDate(date);

	if err != nil {
		return nil, err; 
	}
	
	return lote, nil; 
}