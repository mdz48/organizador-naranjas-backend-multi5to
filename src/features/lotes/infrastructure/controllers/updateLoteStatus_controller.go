package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateLoteStatusController struct {
	uc *application.UpdateLoteStatusUseCase
}

func NewUpdateLoteStatusController(uc *application.UpdateLoteStatusUseCase) *UpdateLoteStatusController {
	return &UpdateLoteStatusController{
		uc: uc,
	}
}

func (ctr *UpdateLoteStatusController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID inválido"})
		return
	}

	var statusRequest struct {
		Estado string `json:"estado" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&statusRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Datos inválidos: se requiere el campo 'estado'"})
		return
	}

	lote, errUpdate := ctr.uc.Run(int(id), statusRequest.Estado)

	if errUpdate != nil {
		ctx.JSON(500, gin.H{"error": errUpdate.Error()})
		return
	}

	ctx.JSON(200, lote)
}
