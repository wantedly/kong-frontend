package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"alert":     false,
		"error":     false,
		"logged_in": true,
		"message":   "",
	})
}
