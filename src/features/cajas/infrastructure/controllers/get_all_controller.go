package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
)

type GetAllController struct {
	getAllService *application.GetAllUseCase
}

func NewGetAllCajaController(getAllService *application.GetAllUseCase) *GetAllController {
	return &GetAllController{getAllService}
}

func (c *GetAllController) GetAll(ctx *gin.Context) {
	cajas, err := c.getAllService.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error al obtener las cajas"})
		return
	}
	ctx.JSON(200, cajas)
}
