package application

import "organizador-naranjas-backend-multi5to/src/features/lotes/domain"

type UpdateLoteUseCase struct {
	loteRepository domain.ILote
}

func NewUpdateLoteUseCase(loteRepository domain.ILote) *UpdateLoteUseCase {
	return &UpdateLoteUseCase{
		loteRepository: loteRepository,
	}
}

func (uc *UpdateLoteUseCase) Run(id int, lote domain.Lote) (domain.Lote, error) {
	lote, err := uc.loteRepository.Update(id, lote)

	if err != nil {
		return domain.Lote{}, err
	}

	return lote, nil
}
