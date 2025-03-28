package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type CajasRoutes struct {
	engine                       *gin.Engine
	getAllController             *controllers.GetAllController
	createCajaController         *controllers.CreateCajaController
	updateCajaController         *controllers.UpdateController
	deleteCajaController         *controllers.DeleteCajaController
	getTop3CajasByLoteController *controllers.GetTop3CajasByLoteController
}

func NewCajasRoutes(
	engine *gin.Engine,
	getAllController *controllers.GetAllController,
	createCajaController *controllers.CreateCajaController,
	updateCajaController *controllers.UpdateController,
	deleteCajaController *controllers.DeleteCajaController,
	getTop3CajasByLoteController *controllers.GetTop3CajasByLoteController,
) *CajasRoutes {
	return &CajasRoutes{
		engine:                       engine,
		getAllController:             getAllController,
		createCajaController:         createCajaController,
		updateCajaController:         updateCajaController,
		deleteCajaController:         deleteCajaController,
		getTop3CajasByLoteController: getTop3CajasByLoteController,
	}
}

func (r *CajasRoutes) SetupRoutes() {
	cajas := r.engine.Group("/cajas")
	{
		cajas.GET("/", r.getAllController.GetAll)
		cajas.POST("/", r.createCajaController.Create)
		cajas.PUT("/:id", r.updateCajaController.Update)
		cajas.DELETE("/:id", r.deleteCajaController.Delete)
		cajas.GET("/lote/:loteId", r.getTop3CajasByLoteController.Handle)
	}
}

func (r *CajasRoutes) Run() error {
	if err := r.engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}
