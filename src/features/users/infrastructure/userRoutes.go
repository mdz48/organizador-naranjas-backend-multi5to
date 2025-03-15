package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	engine *gin.Engine
	createUserController *controllers.CreateUserController
	logInController *controllers.LogInController
}

func NewUserRoutes(engine *gin.Engine, createUserController *controllers.CreateUserController, logInController *controllers.LogInController) *UserRoutes {
	return &UserRoutes{
		engine: engine,
		createUserController: createUserController,
		logInController: logInController,
	}
}

func (routes *UserRoutes) SetupRoutes() {
	userRoutes := routes.engine.Group("/users")
	{
		userRoutes.POST("/", routes.createUserController.Run)
		userRoutes.POST("/auth/login", routes.logInController.Run)
	}
}

func (routes *UserRoutes) Run() error {
	if err := routes.engine.Run(); err != nil {
		return err;
	}

	return nil; 
}