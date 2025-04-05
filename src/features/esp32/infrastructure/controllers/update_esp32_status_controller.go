package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/application"

	"github.com/gin-gonic/gin"
)

type UpdateEsp32StatusController struct {
	uc *application.UpdateEsp32StatusUseCase
}

func NewUpdateEsp32StatusController(uc *application.UpdateEsp32StatusUseCase) *UpdateEsp32StatusController {
	return &UpdateEsp32StatusController{
		uc: uc,
	}
}

func (c *UpdateEsp32StatusController) Run(ctx *gin.Context) {
	id := ctx.Param("id")

	var requestBody struct {
		Status string `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": "Se requiere el campo 'status'"})
		return
	}

	err := c.uc.Run(id, requestBody.Status)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Status actualizado correctamente"})
}
