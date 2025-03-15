package main

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/core"

	cajasUseCases "organizador-naranjas-backend-multi5to/src/features/cajas/application"
	cajasInfrastructure "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure"
	cajasControllers "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure/controllers"

	naranjasInfrastructure "organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure"
    naranjasUseCases "organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	naranjasControllers "organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure/controllers"

	usersInfrastructure "organizador-naranjas-backend-multi5to/src/features/users/infrastructure"
    usersUseCases "organizador-naranjas-backend-multi5to/src/features/users/application"
	usersControllers "organizador-naranjas-backend-multi5to/src/features/users/infrastructure/controllers"
)

type Dependencies struct {
	engine *gin.Engine
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		engine: gin.Default(),
	}
}

func (d *Dependencies) Run() error {
	database := core.NewDatabase()

	cajasDatabase := cajasInfrastructure.NewMySQL(database.Conn)
	getAllUseCase := cajasUseCases.NewGetAllCajaUseCase(cajasDatabase)
	getAllController := cajasControllers.NewGetAllCajaController(getAllUseCase)
	createCajaUseCase := cajasUseCases.NewCreateCajaUseCase(cajasDatabase)
	createCajaController := cajasControllers.NewCreateCajaController(createCajaUseCase)
	cajasRoutes := cajasInfrastructure.NewCajasRoutes(d.engine, getAllController, createCajaController)

	naranjaDatabase := naranjasInfrastructure.NewMySQL(database.Conn); 
	createNaranjaUseCase := naranjasUseCases.NewCreateCajaUseCase(naranjaDatabase)
	createNaranjaController := naranjasControllers.NewCreateNaranjaController(createNaranjaUseCase)
	getAllNaranjaUseCase := naranjasUseCases.NewGetAllUseCase(naranjaDatabase)
	getAllNaranjasController := naranjasControllers.NewGetAllController(getAllNaranjaUseCase)
	updateNaranjaUseCase := naranjasUseCases.NewUpdateNaranjaUseCase(naranjaDatabase)
	updateContollers := naranjasControllers.NewUpdateNaranjaController(updateNaranjaUseCase)
	naranjasRoutes := naranjasInfrastructure.NewNaranjasRoutes(d.engine,createNaranjaController, getAllNaranjasController, updateContollers)

	userDataBase := usersInfrastructure.NewMysql(database.Conn);
	createUser := usersUseCases.NewSaveUser(userDataBase); 
	logInUser := usersUseCases.NewLogInUseCase(userDataBase)
	createUserController := usersControllers.NewCreateUserController(createUser); 
	logInController := usersControllers.NewLoginController(logInUser);
	userRoutes := usersInfrastructure.NewUserRoutes(d.engine, createUserController, logInController); 

	cajasRoutes.SetupRoutes()
	naranjasRoutes.SetupRoutes();
	userRoutes.SetupRoutes(); 

	return d.engine.Run(":8080")
}
