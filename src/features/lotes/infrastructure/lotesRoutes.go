package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type LotesRoutes struct {
	engine *gin.Engine
	createLoteController *controllers.CreateLoteController
	listLotesController *controllers.GetAllLotesController
	listLoteIdController *controllers.ListLoteIdController
	listLoteDateController *controllers.ListLoteDateController
	deleteLoteController *controllers.DeleteLoteController
	updateLoteController *controllers.UpdateLoteController
}

func NewLotesRoutes(
	engine *gin.Engine, 
	createController *controllers.CreateLoteController, 
	listLotesController *controllers.GetAllLotesController,
	listLoteIdControler *controllers.ListLoteIdController,
	listLoteDateController *controllers.ListLoteDateController,
	deleteLoteController *controllers.DeleteLoteController,
	updateLoteController *controllers.UpdateLoteController,
	) *LotesRoutes {
	return &LotesRoutes{
		engine: engine,
		createLoteController: createController,
		listLotesController: listLotesController,
		listLoteIdController: listLoteIdControler,
		listLoteDateController: listLoteDateController,
		deleteLoteController: deleteLoteController,
		updateLoteController: updateLoteController,
	}
}

func (routes *LotesRoutes) SetupRoutes() {
	lotes := routes.engine.Group("/lotes")
	{
		lotes.GET("/", routes.listLotesController.Run)
		lotes.POST("/", routes.createLoteController.Run)
		lotes.GET("/date/:date", routes.listLoteDateController.Run)
		lotes.GET("/:id", routes.listLoteIdController.Run)
		lotes.DELETE("/:id", routes.deleteLoteController.Run)
		lotes.PUT("/:id", routes.updateLoteController.Run)
	}
}

func (routes *LotesRoutes) Run() error {
	if err := routes.engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}
