package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"
)

type CreateNaranjaController struct {
	createNaranjaService *application.CreateNaranjaUseCase
}

func NewCreateNaranjaController(createNaranjaService *application.CreateNaranjaUseCase) *CreateNaranjaController {
	return &CreateNaranjaController{createNaranjaService}
}

func (c *CreateNaranjaController) Create(ctx *gin.Context) {
	var caja domain.Naranja
	if err := ctx.ShouldBindJSON(&caja); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	naranjaCreada, err := c.createNaranjaService.Execute(caja)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error creating naranja"})
		return
	}
	ctx.JSON(201, naranjaCreada)
}
