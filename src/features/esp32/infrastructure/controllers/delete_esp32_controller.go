package controllers

import (
	"organizador-naranjas-backend-multi5to/src/features/esp32/application"
	"github.com/gin-gonic/gin"
)

type DeleteEsp32Controller struct {
	deleteEsp32UseCase *application.DeleteEsp32UseCase
}

func NewDeleteEsp32Controller(deleteEsp32UseCase *application.DeleteEsp32UseCase) *DeleteEsp32Controller {
	return &DeleteEsp32Controller{
		deleteEsp32UseCase: deleteEsp32UseCase,
	}
}

func (controller *DeleteEsp32Controller) Run(c *gin.Context) {
	id := c.Param("id")
	err := controller.deleteEsp32UseCase.Run(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "deleted",
	})
}