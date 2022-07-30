package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HealthCheck (c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "server can be reached",
	})
}
