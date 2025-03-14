package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type UpdateController struct {
	updateService *application.UpdateCajaUseCase
}

func NewUpdateCajaController(updateCaja *application.UpdateCajaUseCase) *UpdateController {
	return &UpdateController{updateCaja}
}

func (c *UpdateController) Update(ctx *gin.Context) {
	var caja domain.Caja
	if err := ctx.ShouldBindJSON(&caja); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	cajaEditada, err := c.updateService.Execute(caja)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error updating caja"})
		return
	}
	ctx.JSON(201, cajaEditada)
}
