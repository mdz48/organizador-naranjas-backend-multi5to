package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/application"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEsp32ByPropietarioAndStatusController struct {
	uc *application.GetEsp32ByPropietarioAndStatusUseCase
}

func NewGetEsp32ByPropietarioAndStatusController(uc *application.GetEsp32ByPropietarioAndStatusUseCase) *GetEsp32ByPropietarioAndStatusController {
	return &GetEsp32ByPropietarioAndStatusController{
		uc: uc,
	}
}

func (c *GetEsp32ByPropietarioAndStatusController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Status fijo como "esperando"
	status := "esperando"

	esp32Devices, err := c.uc.Run(id, status)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Si no hay dispositivos, devolver una lista vacía en lugar de error
	if len(esp32Devices) == 0 {
		ctx.JSON(200, []entities.Esp32{})
		return
	}

	ctx.JSON(200, esp32Devices)
}
