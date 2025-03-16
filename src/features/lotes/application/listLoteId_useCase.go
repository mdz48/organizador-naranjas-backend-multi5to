package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type ListLoteIdUseCase struct {
	loteRepository domain.ILote
}

func NewListLoteIdUseCase(loteRepository domain.ILote) *ListLoteIdUseCase {
	return &ListLoteIdUseCase{
		loteRepository: loteRepository,
	}
}

func (uc *ListLoteIdUseCase) Run(id int) (domain.Lote, error) {
	lote, err := uc.loteRepository.GetById(id); 

	if err != nil {
		return domain.Lote{} ,err; 
	}

	return lote, nil; 
}