package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLotesWithCajasByUserDateRangeController struct {
	useCase *application.GetLotesWithCajasByUserDateRangeUseCase
}

func NewGetLotesWithCajasByUserDateRangeController(useCase *application.GetLotesWithCajasByUserDateRangeUseCase) *GetLotesWithCajasByUserDateRangeController {
	return &GetLotesWithCajasByUserDateRangeController{
		useCase: useCase,
	}
}

func (c *GetLotesWithCajasByUserDateRangeController) Run(ctx *gin.Context) {
	// Obtener el ID de usuario
	userIdStr := ctx.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Obtener parámetros de fechas
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	// Validar que se proporcionaron ambas fechas
	if startDate == "" || endDate == "" {
		ctx.JSON(400, gin.H{"error": "Debe proporcionar ambos parámetros: start_date y end_date"})
		return
	}

	// Ejecutar el caso de uso
	response, err := c.useCase.Execute(userId, startDate, endDate)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, response)
}
