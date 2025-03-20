package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
	"organizador-naranjas-backend-multi5to/src/features/cajas/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateController struct {
	updateService *application.UpdateCajaUseCase
}

func NewUpdateCajaController(updateCaja *application.UpdateCajaUseCase) *UpdateController {
	return &UpdateController{updateCaja}
}

func (c *UpdateController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid caja ID"})
		return
	}

	var caja domain.Caja
	if err := ctx.ShouldBindJSON(&caja); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	caja.ID = int(id)
	updatedCaja, errUpdate := c.updateService.Execute(caja)
	if errUpdate != nil {
		ctx.JSON(500, gin.H{"error": errUpdate.Error()})
		return
	}

	ctx.JSON(200, updatedCaja)
}
