package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type GetAllByJefeUseCase struct {
	db domain.IUser
}

func NewGetAllByJefeUseCase(db domain.IUser) *GetAllByJefeUseCase {
	return &GetAllByJefeUseCase{
		db: db,
	}
}

func (uc *GetAllByJefeUseCase) Run(id_jefe int32) ([]*entities.UserResponse, error) {
	users, err := uc.db.GetAllByJefe(id_jefe)
	if err != nil {
		return nil, err
	}
	
	result := make([]*entities.UserResponse, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	
	return result, nil
}



