package controllers

import (
	"github.com/gin-gonic/gin"
	"organizador-naranjas-backend-multi5to/src/features/cajas/application"
	"strconv"
)

type DeleteCajaController struct {
	deleteService *application.DeleteCajaUseCase
}

func NewDeleteCajaController(deleteCaja *application.DeleteCajaUseCase) *DeleteCajaController {
	return &DeleteCajaController{deleteCaja}
}

func (c *DeleteCajaController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid caja ID"})
		return
	}

	errDelete := c.deleteService.Execute(int(id))
	if errDelete != nil {
		ctx.JSON(500, gin.H{"error": errDelete.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Caja eliminada"})
}