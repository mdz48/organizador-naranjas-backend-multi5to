package application

import (
	"errors"
	"fmt"
	"organizador-naranjas-backend-multi5to/src/core/middlewares"
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

func (uc *LogInUseCase) Run(userLog *entities.UserLogIn) (*entities.Claims, error) {
	fmt.Printf("user: %v\n", userLog)

	// Obtener el usuario de la base de datos
	user, err := uc.userRepository.LogIn(userLog)
	if err != nil {
		return nil, err
	}

	// Verificar la contraseña
	err = middlewares.VerifyPassword(userLog.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Crear los claims con TODOS los campos del usuario
	claims := &entities.Claims{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Rol:      user.Rol,
		Email:    user.Email,
		Id_jefe:  user.Id_jefe,
	}

	return claims, nil
}
