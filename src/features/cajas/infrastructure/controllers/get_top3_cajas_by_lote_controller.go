package controllers

import (
	"strconv"

	"organizador-naranjas-backend-multi5to/src/features/cajas/application"

	"github.com/gin-gonic/gin"
)

type GetTop3CajasByLoteController struct {
	getTop3CajasByLoteUseCase *application.GetTop3CajasByLoteUseCase
}

func NewGetTop3CajasByLoteController(getTop3CajasByLoteUseCase *application.GetTop3CajasByLoteUseCase) *GetTop3CajasByLoteController {
	return &GetTop3CajasByLoteController{getTop3CajasByLoteUseCase: getTop3CajasByLoteUseCase}
}

func (c *GetTop3CajasByLoteController) Handle(ctx *gin.Context) {
	loteIdStr := ctx.Param("loteId")

	loteId, err := strconv.Atoi(loteIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de lote inválido"})
		return
	}

	cajas, err := c.getTop3CajasByLoteUseCase.Execute(loteId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, cajas)
}
