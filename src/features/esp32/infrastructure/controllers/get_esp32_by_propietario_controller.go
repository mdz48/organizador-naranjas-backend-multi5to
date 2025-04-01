package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/application"
	"organizador-naranjas-backend-multi5to/src/features/esp32/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
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
    
    esp32Devices, err := c.uc.Run(id)
    if err != nil {
        ctx.JSON(404, gin.H{"error": err.Error()})
        return
    }
    
    // Si no hay dispositivos, devolver una lista vacía en lugar de error
    if len(esp32Devices) == 0 {
        ctx.JSON(200, []entities.Esp32{})
        return
    }
    
    ctx.JSON(200, esp32Devices)
}