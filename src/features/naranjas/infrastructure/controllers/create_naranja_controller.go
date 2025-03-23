package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/domain"

	"github.com/gin-gonic/gin"
)

type CreateNaranjaController struct {
	createNaranjaService *application.CreateNaranjaUseCase
}

func NewCreateNaranjaController(createNaranjaService *application.CreateNaranjaUseCase) *CreateNaranjaController {
	return &CreateNaranjaController{createNaranjaService}
}

func (c *CreateNaranjaController) Create(ctx *gin.Context) {
	var naranja domain.Naranja 
	if err := ctx.ShouldBindJSON(&naranja); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	naranjaCreada, err := c.createNaranjaService.Execute(naranja)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()}) 
		return
	}
	ctx.JSON(201, naranjaCreada)
}
