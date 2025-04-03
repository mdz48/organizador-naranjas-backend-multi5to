package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type LotesRoutes struct {
	engine                                     *gin.Engine
	createLoteController                       *controllers.CreateLoteController
	listLotesController                        *controllers.GetAllLotesController
	listLoteIdController                       *controllers.ListLoteIdController
	listLoteDateController                     *controllers.ListLoteDateController
	deleteLoteController                       *controllers.DeleteLoteController
	updateLoteController                       *controllers.UpdateLoteController
	updateLoteStatusController                 *controllers.UpdateLoteStatusController
	getLoteByIdController                      *controllers.GetLotesByUserController
	createLoteWithCajas                        *controllers.CreateLoteWithCajasController
	getLoteWithCajasController                 *controllers.GetLoteWithCajasController
	getAllLotesWithCajasController             *controllers.GetAllLotesWithCajasController
	getLotesWithCajasByUserController          *controllers.GetLotesWithCajasByUserController
	getLotesWithCajasByUserDateRangeController *controllers.GetLotesWithCajasByUserDateRangeController
}

func NewLotesRoutes(
	engine *gin.Engine,
	createController *controllers.CreateLoteController,
	listLotesController *controllers.GetAllLotesController,
	listLoteIdControler *controllers.ListLoteIdController,
	listLoteDateController *controllers.ListLoteDateController,
	deleteLoteController *controllers.DeleteLoteController,
	updateLoteController *controllers.UpdateLoteController,
	updateLoteStatusController *controllers.UpdateLoteStatusController,
	getLoteByIdController *controllers.GetLotesByUserController,
	createLoteWithCajas *controllers.CreateLoteWithCajasController,
	getLoteWithCajasController *controllers.GetLoteWithCajasController,
	getAllLotesWithCajasController *controllers.GetAllLotesWithCajasController,
	getLotesWithCajasByUserController *controllers.GetLotesWithCajasByUserController,
	getLotesWithCajasByUserDateRangeController *controllers.GetLotesWithCajasByUserDateRangeController,
) *LotesRoutes {
	return &LotesRoutes{
		engine:                                     engine,
		createLoteController:                       createController,
		listLotesController:                        listLotesController,
		listLoteIdController:                       listLoteIdControler,
		listLoteDateController:                     listLoteDateController,
		deleteLoteController:                       deleteLoteController,
		updateLoteController:                       updateLoteController,
		updateLoteStatusController:                 updateLoteStatusController,
		getLoteByIdController:                      getLoteByIdController,
		createLoteWithCajas:                        createLoteWithCajas,
		getLoteWithCajasController:                 getLoteWithCajasController,
		getAllLotesWithCajasController:             getAllLotesWithCajasController,
		getLotesWithCajasByUserController:          getLotesWithCajasByUserController,
		getLotesWithCajasByUserDateRangeController: getLotesWithCajasByUserDateRangeController,
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
		lotes.PATCH("/:id/status", routes.updateLoteStatusController.Run)
		lotes.GET("/user/:id", routes.getLoteByIdController.Run)
		lotes.POST("/with-cajas", routes.createLoteWithCajas.Run)
		lotes.GET("/with-cajas/:id", routes.getLoteWithCajasController.Run)
		lotes.GET("/with-cajas", routes.getAllLotesWithCajasController.Run)
		lotes.GET("/user/:id/with-cajas", routes.getLotesWithCajasByUserController.Run)
		lotes.GET("/user/:id/with-cajas/date-range", routes.getLotesWithCajasByUserDateRangeController.Run)
	}
}

func (routes *LotesRoutes) Run() error {
	if err := routes.engine.Run(); err != nil {
		return err
	}
	return nil
}
