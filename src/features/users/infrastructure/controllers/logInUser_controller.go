package controllers

import (
	"fmt"
	"log"
	"organizador-naranjas-backend-multi5to/src/core/middlewares"
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type LogInController struct {
	uc *application.LogInUseCase
}

func NewLoginController(uc *application.LogInUseCase) *LogInController {
	return &LogInController{
		uc: uc,
	}
}

func (ctr *LogInController) Run(ctx *gin.Context) {
	var userLog *entities.UserLogIn

	if err := ctx.ShouldBindJSON(&userLog); err != nil {
		ctx.JSON(400, err)
		return
	}
	user, err := ctr.uc.Run(userLog)
	fmt.Printf("user: %s", userLog)
	if err != nil {
		ctx.JSON(401, err)
		return
	}
	errCompare := middlewares.VerifyPassword(userLog.Password, user.Password)
	if errCompare != nil {
		ctx.JSON(401, err)
		return
	}
	token, errToken := middlewares.GenerateToken(user)
	if errToken != nil {
		log.Printf("error %s", errToken)
		ctx.JSON(400, errToken)
		return
	}

	ctx.JSON(201, gin.H{
		"token": token,
	})
}
