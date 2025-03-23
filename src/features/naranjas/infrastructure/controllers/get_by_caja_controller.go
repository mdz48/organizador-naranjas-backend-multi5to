package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/naranjas/application"
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetByCajaNaranjaController struct {
	getByCajaNaranjaUseCase *application.GetByCajaNaranjaUseCase
}

func NewGetByCajaNaranjaController(getByCajaNaranjaUseCase *application.GetByCajaNaranjaUseCase) *GetByCajaNaranjaController {
	return &GetByCajaNaranjaController{getByCajaNaranjaUseCase: getByCajaNaranjaUseCase}
}
func (controller *GetByCajaNaranjaController) Handle(c *gin.Context) {
	cajaIDStr := c.Param("id")
	
	cajaID, err := strconv.Atoi(cajaIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid caja ID"})
		return
	}
	
	naranjas, err := controller.getByCajaNaranjaUseCase.Execute(cajaID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, naranjas)
}
