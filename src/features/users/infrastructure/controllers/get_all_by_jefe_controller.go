package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/users/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetAllByJefeController struct {
	uc *application.GetAllByJefeUseCase
}

func NewGetAllByJefeController(uc *application.GetAllByJefeUseCase) *GetAllByJefeController {
	return &GetAllByJefeController{uc: uc}
}

func (controller *GetAllByJefeController) Run(c *gin.Context) {
	jefeId := c.Param("jefeId")
	jefeIdInt, _ := strconv.Atoi(jefeId)
	users, err := controller.uc.Run(int32(jefeIdInt))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}