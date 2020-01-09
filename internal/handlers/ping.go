package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping is simple keep-alive/ping handler
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
