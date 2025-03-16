package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type RemoveLoteUseCase struct {
	loteRepository domain.ILote
}

func NewRemoveLoteUseCase(loteRepository domain.ILote) *RemoveLoteUseCase {
	return &RemoveLoteUseCase{
		loteRepository: loteRepository,
	}
}

func (uc *RemoveLoteUseCase) Run(id int) error {
	err := uc.loteRepository.Delete(id);

	if err != nil {
		return err; 
	}

	return nil; 
}