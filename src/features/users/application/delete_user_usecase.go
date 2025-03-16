package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type DeleteUserUseCase struct {
	db domain.IUser
}

func NewDeleteUserUseCase(db domain.IUser) *DeleteUserUseCase { return &DeleteUserUseCase{db: db} }

func (uc *DeleteUserUseCase) Run(user *entities.User) (*entities.User, error) {
	user, err := uc.db.Delete(user)

	if err != nil {
		return &entities.User{}, err
	}

	return user, err
}
