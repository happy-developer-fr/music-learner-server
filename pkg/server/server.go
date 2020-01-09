package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/read-music-learner/music-learner-server.git/internal/handlers"
	"log"
)

var (
	PORT string
	HOST string
)

func init() {
	HOST = "localhost"
	PORT = "7777"
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT)
	log.Fatalln(r.Run(HOST + ":" + PORT))
}
