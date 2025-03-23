package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type NaranjasRoutes struct {
	engine                  *gin.Engine
	createNaranjaController *controllers.CreateNaranjaController
	getNaranjaController    *controllers.GetAllNaranjaController
	updateContollers        *controllers.UpdateNaranjaController
	getByEsp32NaranjaController *controllers.GetByEsp32NaranjaController
	getByCajaNaranjaController *controllers.GetByCajaNaranjaController
}

func NewNaranjasRoutes(engine *gin.Engine, createNaranjaController *controllers.CreateNaranjaController, getNaranjaController *controllers.GetAllNaranjaController, updateContollers *controllers.UpdateNaranjaController, getByEsp32NaranjaController *controllers.GetByEsp32NaranjaController, getByCajaNaranjaController *controllers.GetByCajaNaranjaController) *NaranjasRoutes {
	return &NaranjasRoutes{engine: engine, createNaranjaController: createNaranjaController, getNaranjaController: getNaranjaController, updateContollers: updateContollers, getByEsp32NaranjaController: getByEsp32NaranjaController, getByCajaNaranjaController: getByCajaNaranjaController}
}

func (r *NaranjasRoutes) SetupRoutes() {
	naranjas := r.engine.Group("/naranjas")
	{
		naranjas.GET("/", r.getNaranjaController.GetAll)
		naranjas.POST("/", r.createNaranjaController.Create)
		naranjas.PUT("/:id", r.updateContollers.Update)
		naranjas.GET("/esp32/:esp32Id", r.getByEsp32NaranjaController.GetByEsp32)
		naranjas.GET("/caja/:id", r.getByCajaNaranjaController.Handle)
	}
}

func (r *NaranjasRoutes) Run() error {
	if err := r.engine.Run(); err != nil {
		return err
	}
	return nil
}