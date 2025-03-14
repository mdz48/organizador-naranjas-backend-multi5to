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
}

func NewNaranjasRoutes(engine *gin.Engine, createNaranjaController *controllers.CreateNaranjaController, getNaranjaController *controllers.GetAllNaranjaController, updateContollers *controllers.UpdateNaranjaController) *NaranjasRoutes {
	return &NaranjasRoutes{engine: engine, createNaranjaController: createNaranjaController, getNaranjaController: getNaranjaController, updateContollers: updateContollers}
}

func (r *NaranjasRoutes) SetupRoutes() {
	naranjas := r.engine.Group("/naranjas")
	{
		naranjas.GET("/", r.getNaranjaController.GetAll)
		naranjas.POST("/", r.createNaranjaController.Create)
		naranjas.PUT("/:id", r.updateContollers.Update)
	}
}

func (r *NaranjasRoutes) Run() error {
	if err := r.engine.Run(); err != nil {
		return err
	}
	return nil
}