package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"

	"github.com/gin-gonic/gin"
)

type GetAllLotesController struct {
	uc *application.ListLotesUseCase
}

func NewGetAllLotesController(uc *application.ListLotesUseCase) *GetAllLotesController {
	return &GetAllLotesController{
		uc: uc,
	}
}

func (ctr *GetAllLotesController) Run(ctx *gin.Context) {
	lotes, err := ctr.uc.Run();

	if err != nil {
		ctx.JSON(500, err);
		return;
	}

	ctx.JSON(201, lotes);
}