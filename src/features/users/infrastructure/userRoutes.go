package infrastructure

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/users/infrastructure/controllers"
)

type UserRoutes struct {
	engine               *gin.Engine
	createUserController *controllers.CreateUserController
	logInController      *controllers.LogInController
	updateUserController *controllers.UpdateUserController
	deleteUserController *controllers.DeleteUserController
}

func NewUserRoutes(engine *gin.Engine, createUserController *controllers.CreateUserController, logInController *controllers.LogInController, updateUserController *controllers.UpdateUserController, deleteUserController *controllers.DeleteUserController) *UserRoutes {
	return &UserRoutes{
		engine:               engine,
		createUserController: createUserController,
		logInController:      logInController,
		updateUserController: updateUserController,
		deleteUserController: deleteUserController,
	}
}

func (routes *UserRoutes) SetupRoutes() {
	userRoutes := routes.engine.Group("/users")
	{
		userRoutes.POST("/", routes.createUserController.Run)
		userRoutes.POST("/login", routes.logInController.Run)
		userRoutes.PUT("/:id", routes.updateUserController.Run)
		userRoutes.DELETE("/:id", routes.deleteUserController.Run)
	}
}

func (routes *UserRoutes) Run() error {
	if err := routes.engine.Run(); err != nil {
		return err
	}

	return nil
}
