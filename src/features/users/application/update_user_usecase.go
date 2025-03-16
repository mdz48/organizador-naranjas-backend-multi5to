package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type UpdateUserUseCase struct {
	db domain.IUser
}

func NewUpdateUserUseCase(db domain.IUser) *UpdateUserUseCase { return &UpdateUserUseCase{db: db} }

func (uc *UpdateUserUseCase) Run(user *entities.User) (*entities.User, error) {
	user, err := uc.db.Update(user)

	if err != nil {
		return &entities.User{}, err
	}

	return user, err
}
