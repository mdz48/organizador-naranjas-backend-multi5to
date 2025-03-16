package controllers

import (
	"github.com/gin-gonic/gin"
	_ "organizador-naranjas-backend-multi5to/src/core/middlewares"
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"
	"strconv"
)

type DeleteUserController struct {
	uc *application.DeleteUserUseCase
}

func NewDeleteUserController(uc *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		uc: uc,
	}
}

func (ctr *DeleteUserController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	var user entities.User
	user.ID = int32(id)
	deletedUser, errDelete := ctr.uc.Run(&user)
	if errDelete != nil {
		ctx.JSON(500, gin.H{"error": errDelete.Error()})
		return
	}
	ctx.JSON(200, deletedUser)
}
