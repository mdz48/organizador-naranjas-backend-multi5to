package infrastructure

import (
	"organizador-naranjas-backend-multi5to/src/features/sp32/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type Sp32Routes struct {
	engine gin.RouterGroup
	createSp32Controller *controllers.CreateSpr32Controller
}

func NewSp32Routes(engine gin.RouterGroup, createSp32Controller *controllers.CreateSpr32Controller) *Sp32Routes {
	return &Sp32Routes{
		engine: engine,
		createSp32Controller: createSp32Controller,
	}	
}

func (routes *Sp32Routes) Run() {
	sp32Routes := routes.engine.Group("/sp32")
	{
		sp32Routes.POST("/", routes.createSp32Controller.Run)
	}	
}