package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLotesWithCajasByUserController struct {
	getLotesWithCajasByUserUseCase *application.GetLotesWithCajasByUserUseCase
}

func NewGetLotesWithCajasByUserController(getLotesWithCajasByUserUseCase *application.GetLotesWithCajasByUserUseCase) *GetLotesWithCajasByUserController {
	return &GetLotesWithCajasByUserController{
		getLotesWithCajasByUserUseCase: getLotesWithCajasByUserUseCase,
	}
}

func (c *GetLotesWithCajasByUserController) Run(ctx *gin.Context) {
	// Obtener el ID de usuario de los parámetros de la URL
	userIdStr := ctx.Param("id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Ejecutar el caso de uso
	response, err := c.getLotesWithCajasByUserUseCase.Execute(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Devolver la respuesta
	ctx.JSON(200, response)
}
