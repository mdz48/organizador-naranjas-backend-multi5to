package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"github.com/gin-gonic/gin"
)

type GetUsersController struct {
	uc *application.GetUsersUseCase
}

func NewGetUsersController(uc *application.GetUsersUseCase) *GetUsersController {
	return &GetUsersController{
		uc: uc,
	}
}

func (ctr *GetUsersController) Run(ctx *gin.Context) {
	users, err := ctr.uc.GetAll()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, users)
}