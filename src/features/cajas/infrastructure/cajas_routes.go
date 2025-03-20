package infrastructure

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure/controllers"
)

type CajasRoutes struct {
	engine               *gin.Engine
	getAllController     *controllers.GetAllController
	createCajaController *controllers.CreateCajaController
	updateCajaController *controllers.UpdateController
	deleteCajaController *controllers.DeleteCajaController
}

func NewCajasRoutes(engine *gin.Engine, getAllController *controllers.GetAllController, createCajaController *controllers.CreateCajaController, updateCajaController *controllers.UpdateController, deleteCajaController *controllers.DeleteCajaController) *CajasRoutes {
	return &CajasRoutes{
		engine:               engine,
		getAllController:     getAllController,
		createCajaController: createCajaController,
		updateCajaController: updateCajaController,
		deleteCajaController: deleteCajaController,
	}
}

func (r *CajasRoutes) SetupRoutes() {
	cajas := r.engine.Group("/cajas")
	{
		cajas.GET("/", r.getAllController.GetAll)
		cajas.POST("/", r.createCajaController.Create)
		cajas.PUT("/:id", r.updateCajaController.Update)
		cajas.DELETE("/:id", r.deleteCajaController.Delete)
	}
}

func (r *CajasRoutes) Run() error {
	if err := r.engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}
