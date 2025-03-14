package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
)

type GetAllNaranjaController struct {
	getAllService *application.GetAllNaranjaUseCase
}

func NewGetAllController(getAllService *application.GetAllNaranjaUseCase) *GetAllNaranjaController {
	return &GetAllNaranjaController{getAllService}
}

func (c *GetAllNaranjaController) GetAll(ctx *gin.Context) {
	naranjas, err := c.getAllService.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, naranjas)
}