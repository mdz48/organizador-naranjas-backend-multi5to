package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateLoteController struct {
	uc *application.CreateLoteUseCase
}

func NewCreateLoteController(uc *application.CreateLoteUseCase) *CreateLoteController {
	return &CreateLoteController{
		uc: uc,
	}
}

func (ctr *CreateLoteController) Run(ctx *gin.Context) {
	var lote domain.Lote

	if err := ctx.ShouldBindJSON(&lote); err != nil {
		ctx.JSON(400, err)
		return
	}

	if lote.Estado == "" {
		lote.Estado = "cargando"
	}

	// Formatear la fecha correctamente si es un formato ISO
	if strings.Contains(lote.Fecha, "T") {
		parsedTime, err := time.Parse(time.RFC3339, lote.Fecha)
		if err == nil {
			lote.Fecha = parsedTime.Format("2006-01-02")
		}
	}

	lote, err := ctr.uc.Execute(lote)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, lote)
}
