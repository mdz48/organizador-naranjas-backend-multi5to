package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"

	"github.com/gin-gonic/gin"
)

type ListLoteDateController struct {
	uc *application.ListLoteDateUseCase
}

func NewListLoteDateController(uc *application.ListLoteDateUseCase) *ListLoteDateController {
	return &ListLoteDateController{
		uc: uc,
	}
}

func (ctr *ListLoteDateController) Run(ctx *gin.Context) {
	date := ctx.Param("date")

	lote, err := ctr.uc.Run(date);

	if err != nil {
		ctx.JSON(500, err); 
		return; 
	}

	ctx.JSON(201, lote); 
}