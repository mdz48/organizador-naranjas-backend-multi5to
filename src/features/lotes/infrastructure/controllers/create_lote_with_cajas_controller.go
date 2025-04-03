// Nuevo archivo: src/features/lotes/infrastructure/controllers/create_lote_with_cajas_controller.go
package controllers

import (
    "organizador-naranjas-backend-multi5to/src/features/lotes/application"
    "organizador-naranjas-backend-multi5to/src/features/lotes/domain"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

type CreateLoteWithCajasController struct {
    uc *application.CreateLoteWithCajasUseCase
}

func NewCreateLoteWithCajasController(uc *application.CreateLoteWithCajasUseCase) *CreateLoteWithCajasController {
    return &CreateLoteWithCajasController{
        uc: uc,
    }
}

func (ctr *CreateLoteWithCajasController) Run(ctx *gin.Context) {
    var req domain.LoteWithCajas

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Formatear la fecha si es necesario
    if strings.Contains(req.Lote.Fecha, "T") {
        parsedTime, err := time.Parse(time.RFC3339, req.Lote.Fecha)
        if err == nil {
            req.Lote.Fecha = parsedTime.Format("2006-01-02")
        }
    }

    // Llamar al caso de uso con el esp32_id
    createdLote, cajas, err := ctr.uc.Execute(req.Lote, req.Esp32FK)
    if err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(201, gin.H{
        "lote": createdLote,
        "cajas": cajas,
    })
}