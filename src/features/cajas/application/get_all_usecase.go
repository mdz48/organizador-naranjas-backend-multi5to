package application

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type GetAllUseCase struct {
	db domain.ICaja
}

func NewGetAllUseCase(db domain.ICaja) *GetAllUseCase { return &GetAllUseCase{db: db} }

func (g *GetAllUseCase) Execute() ([]domain.Caja, error) {
	cajas, err := g.db.GetAll()
	if err != nil {
		return nil, err
	}
	return cajas, nil
}
