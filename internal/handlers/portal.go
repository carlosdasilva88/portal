package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPortal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Portal online",
	})
}
