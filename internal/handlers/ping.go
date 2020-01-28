package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type reponse struct {
	Message string `json:"message"`
}

// PingWithName is simple keep-alive/ping handler
func PingWithName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		resp := reponse{Message: "Hello " + name}
		log.Printf("PingWithName %v ", resp)
		c.JSON(http.StatusOK, resp)
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := reponse{Message: "Pong"}
		log.Printf("PingWithName %v ", resp)
		c.JSON(http.StatusOK, resp)
	}
}
