package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/lotes/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLotesByUserController struct {
	uc *application.GetLotesByUserUseCase
}

// NewGetLotesByUserUseCase creates a new instance of GetLotesByUserUseCase
func NewGetLotesByUserController(uc *application.GetLotesByUserUseCase) *GetLotesByUserController {
	return &GetLotesByUserController{
		uc: uc,
	}
}

// Execute handles the request to get lots by user ID
func (g *GetLotesByUserController) Run(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	lotes, err := g.uc.Execute(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve lots"})
		return
	}
	// response := make([]map[string]interface{}, len(lotes))
	// for i, lote := range lotes {
	// 	response[i] = map[string]interface{}{
	// 		"id":          lote.ID,
	// 		"estado":      lote.Estado,
	// 		"observaciones": lote.Observaciones,
	// 		"user_id":  lote.UserID,
	// 	}
	// }
	c.JSON(200, lotes)
}