package application

import (
	"organizador-naranjas-backend-multi5to/src/features/users/domain"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
    "organizador-naranjas-backend-multi5to/src/core/middlewares"
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
    // Hashear la contraseña antes de guardarla
    hashedPassword, err := middlewares.HashPassword(user.Password)
    if err != nil {
        return nil, err
    }
    
    // Reemplazar la contraseña en texto plano con la versión hasheada
    user.Password = hashedPassword
    
    // Guardar el usuario con la contraseña hasheada
    return uc.userRepository.Save(user)
}