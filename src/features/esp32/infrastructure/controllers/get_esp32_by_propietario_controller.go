package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/esp32/application"
	"strconv"
)

type GetEsp32ByPropietarioController struct {
	uc *application.GetEsp32ByPropietarioUseCase
}

func NewGetEsp32ByPropietarioController(uc *application.GetEsp32ByPropietarioUseCase) *GetEsp32ByPropietarioController {
	return &GetEsp32ByPropietarioController{
		uc: uc,
	}
}

func (c *GetEsp32ByPropietarioController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id parameter"})
		return
	}
	
	esp32, err := c.uc.Run(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(200, esp32)
}