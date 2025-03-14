package infrastructure

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure/controllers"
)

type CajasRoutes struct {
	engine               *gin.Engine
	createCajaController *controllers.CreateNaranjaController
}

func NewCajasRoutes(engine *gin.Engine, createCajaController *controllers.CreateNaranjaController) *CajasRoutes {
	return &CajasRoutes{
		engine:               engine,
		createCajaController: createCajaController,
	}
}

func (r *CajasRoutes) SetupRoutes() {
	cajas := r.engine.Group("/naranjas")
	{
		cajas.POST("/", r.createCajaController.Create)
	}
}

func (r *CajasRoutes) Run() error {
	if err := r.engine.Run(":8081"); err != nil {
		return err
	}
	return nil
}
