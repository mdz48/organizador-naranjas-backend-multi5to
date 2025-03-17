package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	_"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type UpdateNaranjaController struct {
	updateNaranjaService *application.UpdateNaranjaUseCase
}

func NewUpdateNaranjaController(updateNaranjaService *application.UpdateNaranjaUseCase) *UpdateNaranjaController {
	return &UpdateNaranjaController{updateNaranjaService: updateNaranjaService}
}

func (c *UpdateNaranjaController) Update(ctx *gin.Context) {
	// var caja domain.Naranja
	// if err := ctx.ShouldBindJSON(&caja); err != nil {
	// 	ctx.JSON(400, gin.H{"error": "Invalid input"})
	// 	return
	// }
	// 
	// naranjaCreada, err := c.createNaranjaService.Execute(caja)
	// if err != nil {
	// 	ctx.JSON(500, gin.H{"error": "Error creating naranja"})
	// 	return
	// }
	// ctx.JSON(201, naranjaCreada)
}