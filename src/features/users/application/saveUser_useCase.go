package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type SaveUserUseCase struct {
	userRepository domain.IUser
}

func NewSaveUser(userRepository domain.IUser) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *SaveUserUseCase) Run(user *entities.User) (*entities.UserResponse, error) {
	return uc.userRepository.Save(user)
}