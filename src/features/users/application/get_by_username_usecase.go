package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type GetUserByUsernameUseCase struct {
	userRepository domain.IUser
}

func NewGetUserByUsernameUseCase(userRepository domain.IUser) *GetUserByUsernameUseCase {
	return &GetUserByUsernameUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserByUsernameUseCase) Run(username string) (*entities.UserResponse, error) {
	return uc.userRepository.GetByUsername(username)
}