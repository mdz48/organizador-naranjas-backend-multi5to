package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateLoteController struct {
	uc *application.UpdateLoteUseCase
}

func NewUpdateLoteController(uc *application.UpdateLoteUseCase) *UpdateLoteController {
	return &UpdateLoteController{
		uc: uc,
	}
}

func (ctr *UpdateLoteController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil {
		ctx.JSON(400, err)
		return
	}

	var lote domain.Lote

	if err := ctx.ShouldBindJSON(&lote); err != nil {
		ctx.JSON(400, err)
		return
	}

	// Verificar si el estado está vacío, en ese caso mantenerlo como "cargando"
	if lote.Estado == "" {
		lote.Estado = "cargando"
	}

	// Formatear la fecha correctamente si es un formato ISO
	if strings.Contains(lote.Fecha, "T") {
		parsedTime, err := time.Parse(time.RFC3339, lote.Fecha)
		if err == nil {
			// Convertir a formato YYYY-MM-DD
			lote.Fecha = parsedTime.Format("2006-01-02")
		}
	}

	lote, errUpdate := ctr.uc.Run(int(id), lote)

	if errUpdate != nil {
		ctx.JSON(500, errUpdate)
		return
	}

	ctx.JSON(201, lote)
}
