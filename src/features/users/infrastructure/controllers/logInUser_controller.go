package controllers

import (
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
	var userLog entities.UserLogIn

	if err := ctx.ShouldBindJSON(&userLog); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Ejecutar el caso de uso para obtener los claims
	claims, err := ctr.uc.Run(&userLog)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// Generar el token a partir de los claims
	token, errToken := middlewares.GenerateTokenFromClaims(claims)
	if errToken != nil {
		log.Printf("error: %s", errToken)
		ctx.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	// Retornar token y claims en la respuesta
	ctx.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"name":     claims.Name,
			"username": claims.Username,
			"rol":      claims.Rol,
		},
	})
}
