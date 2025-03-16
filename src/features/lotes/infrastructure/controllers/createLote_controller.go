package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"

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
		return; 
	}

	lote, errCreate := ctr.uc.Execute(lote)

	if errCreate != nil {
		ctx.JSON(500, errCreate)
		return; 
	}

	ctx.JSON(201, lote); 
}