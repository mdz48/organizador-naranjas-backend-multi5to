package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListLoteIdController struct {
	uc *application.ListLoteIdUseCase
}

func NewListLoteIdController(uc *application.ListLoteIdUseCase) *ListLoteIdController {
	return &ListLoteIdController{
		uc: uc,
	}
}

func (ctr *ListLoteIdController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32);

	if err != nil {
		ctx.JSON(400,  err);
		return; 
	}

	lote, errSearch := ctr.uc.Run(int(id));

	if errSearch != nil {
		ctx.JSON(500, errSearch); 
		return; 
	}

	ctx.JSON(201, lote); 
}