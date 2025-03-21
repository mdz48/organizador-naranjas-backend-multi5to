package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetUserByIdController struct {
	uc *application.GetUserByIDUseCase
}

func NewGetUserByIdController(uc *application.GetUserByIDUseCase) *GetUserByIdController {
	return &GetUserByIdController{
		uc: uc,
	}
}
func (ctr *GetUserByIdController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	user, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, user)
}