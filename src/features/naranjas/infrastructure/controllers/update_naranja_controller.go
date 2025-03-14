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
	//naranja := domain.Naranja{}
	//ctx.BindJSON(&naranja)
	//c.updateNaranjaService.Execute(naranja)
	ctx.JSON(200, gin.H{"message": "Naranja updated"})
}