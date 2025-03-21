package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type GetUsersUseCase struct {
	userRepository domain.IUser
}

func NewGetUsersUseCase(userRepository domain.IUser) *GetUsersUseCase {
	return &GetUsersUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUsersUseCase) GetAll() ([]entities.UserResponse, error) {
	return uc.userRepository.GetAll()
}

