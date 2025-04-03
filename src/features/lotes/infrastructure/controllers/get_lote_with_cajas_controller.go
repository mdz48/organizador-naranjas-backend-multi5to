package controllers

import (
    "strconv"
    "organizador-naranjas-backend-multi5to/src/features/lotes/application"
    "github.com/gin-gonic/gin"
)

type GetLoteWithCajasController struct {
    getLoteWithCajasUseCase *application.GetLoteWithCajasUseCase
}

func NewGetLoteWithCajasController(getLoteWithCajasUseCase *application.GetLoteWithCajasUseCase) *GetLoteWithCajasController {
    return &GetLoteWithCajasController{
        getLoteWithCajasUseCase: getLoteWithCajasUseCase,
    }
}

func (c *GetLoteWithCajasController) Run(ctx *gin.Context) {
    // Obtener el ID del lote de los parámetros de la URL
    loteIdStr := ctx.Param("id")
    
    loteId, err := strconv.Atoi(loteIdStr)
    if err != nil {
        ctx.JSON(400, gin.H{"error": "ID de lote inválido"})
        return
    }
    
    // Ejecutar el caso de uso
    response, err := c.getLoteWithCajasUseCase.Execute(loteId)
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    // Devolver la respuesta
    ctx.JSON(200, response)
}