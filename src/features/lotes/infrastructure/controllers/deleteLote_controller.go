package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteLoteController struct {
	uc *application.RemoveLoteUseCase
}

func NewDeleteLoteController(uc *application.RemoveLoteUseCase) *DeleteLoteController {
	return &DeleteLoteController{
		uc: uc,
	}
}

func (ctr *DeleteLoteController) Run(ctx *gin.Context) {
	idParams := ctx.Param("id");
	id, err := strconv.ParseInt(idParams, 10, 32); 

	if err != nil {
		ctx.JSON(400, err); 
		return;
	}

	errResult := ctr.uc.Run(int(id));

	if errResult != nil {
		ctx.JSON(500, errResult)
		return;
	}

	ctx.JSON(201, "lote deleted")
}