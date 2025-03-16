package domain

import "organizador-naranjas-backend-multi5to/src/features/users/domain/entities"

type IUser interface {
	Save(user *entities.User) (*entities.User, error)
	LogIn(userLog *entities.UserLogIn) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(user *entities.User) (*entities.User, error)
}
