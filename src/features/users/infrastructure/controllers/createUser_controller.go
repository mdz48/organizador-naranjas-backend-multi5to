package controllers

import (
	"organizador-naranjas-backend-multi5to/src/core/middlewares"
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"organizador-naranjas-backend-multi5to/src/features/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	uc *application.SaveUserUseCase
}

func NewCreateUserController(uc *application.SaveUserUseCase) *CreateUserController {
	return &CreateUserController{
		uc: uc,
	}
}

func (ctr *CreateUserController) Run(ctx *gin.Context) {
	var user *entities.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, err);
		return; 
	}
	
	hashPassword, errHashPassword := middlewares.HashPassword(user.Password)

	if errHashPassword != nil {
		ctx.JSON(400, errHashPassword); 
		return;
	}

	user.Password = hashPassword; 

	user, errCreate := ctr.uc.Run(user);

	if errCreate != nil {
		ctx.JSON(500, errCreate); 
		return; 
	}

	ctx.JSON(201, user); 
}
