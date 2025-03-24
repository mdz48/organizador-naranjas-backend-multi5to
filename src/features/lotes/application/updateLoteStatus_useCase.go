package application

import (
	cajaDomain "organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	lotedomain "organizador-naranjas-backend-multi5to/src/features/lotes/domain"
)

type UpdateLoteStatusUseCase struct {
	loteRepository lotedomain.ILote
	cajaRepository cajaDomain.ICaja
}

func NewUpdateLoteStatusUseCase(loteRepository lotedomain.ILote, cajaRepository cajaDomain.ICaja) *UpdateLoteStatusUseCase {
	return &UpdateLoteStatusUseCase{
		loteRepository: loteRepository,
		cajaRepository: cajaRepository,
	}
}

func (uc *UpdateLoteStatusUseCase) Run(id int, estado string) (lotedomain.Lote, error) {
	lote, err := uc.loteRepository.GetById(id)
	if err != nil {
		return lotedomain.Lote{}, err
	}
	
	lote.Estado = estado
	updatedLote, err := uc.loteRepository.Update(id, lote)
	if err != nil {
		return lotedomain.Lote{}, err
	}

	if estado == "terminado" {
		err = uc.cajaRepository.UpdateStatusByLoteId(id, estado)
		if err != nil {
			return lotedomain.Lote{}, err
		}
	}

	return updatedLote, nil
}
