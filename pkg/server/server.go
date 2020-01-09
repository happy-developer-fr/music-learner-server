package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/read-music-learner/music-learner-server.git/internal/handlers"
	"gitlab.com/read-music-learner/music-learner-server.git/pkg/utils"
	"log"
)

var (
	host string
	port string
)

func init() {
	host = utils.MustGet("MUSIC_LEARNER_SERVER_HOST")
	port = utils.MustGet("MUSIC_LEARNER_SERVER_PORT")
}

// Run web server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
