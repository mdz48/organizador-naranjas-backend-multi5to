package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/sp32/application"
	"organizador-naranjas-backend-multi5to/src/features/sp32/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateSpr32Controller struct {
	uc *application.SaveSp32UseCase
}

func NewCreateSp32Controller(uc *application.SaveSp32UseCase) *CreateSpr32Controller {
	return &CreateSpr32Controller{
		uc: uc,
	}
}

func (ctr *CreateSpr32Controller) Run(ctx *gin.Context) {
	var sp32 *entities.Sp32

	if err := ctx.ShouldBindJSON(&sp32); err != nil {
		ctx.JSON(400, err)
		return
	}

	sp32, err := ctr.uc.Run(sp32)

	if err != nil {
		ctx.JSON(500, err)
		return 
	}

	ctx.JSON(201, sp32)
}