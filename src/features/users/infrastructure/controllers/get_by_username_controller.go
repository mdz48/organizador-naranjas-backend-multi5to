package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"github.com/gin-gonic/gin"
)

type GetByUsernameController struct {
	uc *application.GetUserByUsernameUseCase
}

func NewGetByUsernameController(uc *application.GetUserByUsernameUseCase) *GetByUsernameController {
	return &GetByUsernameController{
		uc: uc,
	}
}

func (ctr *GetByUsernameController) Run(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := ctr.uc.Run(username)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, user)
}