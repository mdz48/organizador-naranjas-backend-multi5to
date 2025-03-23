package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"

	"github.com/gin-gonic/gin"
)

type GetByEsp32NaranjaController struct {
	getByEsp32NaranjaUseCase *application.GetByEsp32NaranjaUseCase
}

func NewGetByEsp32NaranjaController(getByEsp32NaranjaUseCase *application.GetByEsp32NaranjaUseCase) *GetByEsp32NaranjaController {
	return &GetByEsp32NaranjaController{getByEsp32NaranjaUseCase: getByEsp32NaranjaUseCase}
}

func (c *GetByEsp32NaranjaController) GetByEsp32(ctx *gin.Context) {
	esp32Id := ctx.Param("esp32Id")

	naranjas, err := c.getByEsp32NaranjaUseCase.Execute(esp32Id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, naranjas)
}
