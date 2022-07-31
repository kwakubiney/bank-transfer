package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "server can be reached",
	})
}
