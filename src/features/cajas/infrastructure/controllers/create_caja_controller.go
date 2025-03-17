package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
)

type CreateCajaController struct {
	createCajaService *application.CreateCajaUseCase
}

func NewCreateCajaController(createCajaService *application.CreateCajaUseCase) *CreateCajaController {
	return &CreateCajaController{createCajaService}
}

func (c *CreateCajaController) Create(ctx *gin.Context) {
	var caja domain.Caja
	if err := ctx.ShouldBindJSON(&caja); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	cajaCreada, err := c.createCajaService.Execute(caja)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, cajaCreada)
}
