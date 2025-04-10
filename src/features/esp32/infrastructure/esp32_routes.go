package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type Esp32Routes struct {
	engine                                   *gin.Engine
	createSp32Controller                     *controllers.CreateEsp32Controller
	getSp32ByUsernameController              *controllers.GetEsp32ByPropietarioController
	deleteEsp32Controller                    *controllers.DeleteEsp32Controller
	updateEsp32StatusController              *controllers.UpdateEsp32StatusController
	getEsp32ByPropietarioAndStatusController *controllers.GetEsp32ByPropietarioAndStatusController
}

func NewEsp32Routes(
	engine *gin.Engine,
	createSp32Controller *controllers.CreateEsp32Controller,
	getSp32ByUsernameController *controllers.GetEsp32ByPropietarioController,
	deleteEsp32Controller *controllers.DeleteEsp32Controller,
	updateEsp32StatusController *controllers.UpdateEsp32StatusController,
	getEsp32ByPropietarioAndStatusController *controllers.GetEsp32ByPropietarioAndStatusController,
) *Esp32Routes {
	return &Esp32Routes{
		engine:                                   engine,
		createSp32Controller:                     createSp32Controller,
		getSp32ByUsernameController:              getSp32ByUsernameController,
		deleteEsp32Controller:                    deleteEsp32Controller,
		updateEsp32StatusController:              updateEsp32StatusController,
		getEsp32ByPropietarioAndStatusController: getEsp32ByPropietarioAndStatusController,
	}
}

func (routes *Esp32Routes) Run() {
	sp32Routes := routes.engine.Group("/esp32")
	{
		sp32Routes.POST("/", routes.createSp32Controller.Run)
		sp32Routes.GET("/propietario/:id", routes.getSp32ByUsernameController.Run)
		sp32Routes.DELETE("/:id", routes.deleteEsp32Controller.Run)
		sp32Routes.PATCH("/:id/status", routes.updateEsp32StatusController.Run)
		sp32Routes.GET("/propietario/:id/waiting", routes.getEsp32ByPropietarioAndStatusController.Run)
	}
}
