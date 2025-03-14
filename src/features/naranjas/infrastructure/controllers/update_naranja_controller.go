package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type UpdateNaranjaController struct {
	updateNaranjaService *application.UpdateNaranjaUseCase
}

func NewUpdateNaranjaController(updateNaranjaService *application.UpdateNaranjaUseCase) *UpdateNaranjaController {
	return &UpdateNaranjaController{updateNaranjaService: updateNaranjaService}
}

func (c *UpdateNaranjaController) Update(ctx *gin.Context) {
	var naranja domain.Naranja
	if err := ctx.ShouldBindJSON(&naranja); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	naranjaEditada, err := c.updateNaranjaService.Execute(naranja)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error updating naranja"})
		return
	}
	ctx.JSON(201, naranjaEditada)
}
