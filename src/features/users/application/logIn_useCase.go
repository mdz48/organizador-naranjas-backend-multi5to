package application

import (
	"fmt"
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
)

type LogInUseCase struct {
	userRepository domain.IUser
}

func NewLogInUseCase(userRepository domain.IUser) *LogInUseCase {
	return &LogInUseCase{
		userRepository: userRepository,
	}
}

func (uc *LogInUseCase) Run(userLog *entities.UserLogIn) (*entities.User, error) {
	fmt.Printf("user: %s", userLog);
	loged, errLoged := uc.userRepository.LogIn(userLog);

	if errLoged != nil {
		return &entities.User{}, errLoged; 
	}

	return loged, errLoged; 
}