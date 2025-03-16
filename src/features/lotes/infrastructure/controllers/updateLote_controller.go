package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"organizador-naranjas-backend-multi5to/src/features/lotes/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateLoteController struct {
	uc *application.UpdateLoteUseCase
}

func NewUpdateLoteController(uc *application.UpdateLoteUseCase) *UpdateLoteController {
	return &UpdateLoteController{
		uc: uc,
	}
}

func (ctr *UpdateLoteController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id");
	id, err := strconv.ParseInt(idParam, 10, 32);

	if err != nil {
		ctx.JSON(400, err); 
		return; 
	}	

	var lote domain.Lote

	if err := ctx.ShouldBindJSON(&lote); err != nil {
		ctx.JSON(400, err)
		return; 
	}

	lote, errUpdate := ctr.uc.Run(int(id), lote);

	if errUpdate != nil {
		ctx.JSON(500, errUpdate); 
		return;
	}

	ctx.JSON(201, lote); 
}