package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type GetUserByIDUseCase struct {
	userRepository domain.IUser
}

func NewGetUserByIDUseCase(userRepository domain.IUser) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserByIDUseCase) Run(id int32) (*entities.UserResponse, error) {
	return uc.userRepository.GetByID(id)
}