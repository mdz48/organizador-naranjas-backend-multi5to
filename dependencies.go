package main

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/core"

	cajasUseCases "organizador-naranjas-backend-multi5to/src/features/cajas/application"
	cajasInfrastructure "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure"
	cajasControllers "organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure/controllers"
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
	getAllUseCase := cajasUseCases.NewGetAllUseCase(cajasDatabase)
	getAllController := cajasControllers.NewGetAllController(getAllUseCase)
	createCajaUseCase := cajasUseCases.NewCreateCajaUseCase(cajasDatabase)
	createCajaController := cajasControllers.NewCreateCajaController(createCajaUseCase)
	cajasRoutes := cajasInfrastructure.NewCajasRoutes(d.engine, getAllController, createCajaController)
	cajasRoutes.SetupRoutes()

	return d.engine.Run(":8080")
}
