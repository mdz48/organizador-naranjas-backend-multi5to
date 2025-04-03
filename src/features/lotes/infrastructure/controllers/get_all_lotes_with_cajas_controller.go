package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"

	"github.com/gin-gonic/gin"
)

type GetAllLotesWithCajasController struct {
    getAllLotesWithCajasUseCase *application.GetAllLotesWithCajasUseCase
}

func NewGetAllLotesWithCajasController(getAllLotesWithCajasUseCase *application.GetAllLotesWithCajasUseCase) *GetAllLotesWithCajasController {
    return &GetAllLotesWithCajasController{
        getAllLotesWithCajasUseCase: getAllLotesWithCajasUseCase,
    }
}

func (c *GetAllLotesWithCajasController) Run(ctx *gin.Context) {
    // Ejecutar el caso de uso para obtener todos los lotes con cajas
    response, err := c.getAllLotesWithCajasUseCase.Execute()
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    // Devolver la respuesta
    ctx.JSON(200, response)
}