package server

import (
	"github.com/gin-gonic/gin"
	"github.com/happy-developer-fr/music-learner-server/internal/db"
	"github.com/happy-developer-fr/music-learner-server/internal/handlers"
	"github.com/happy-developer-fr/music-learner-server/pkg/utils"
	"log"
)

var (
	port string
)

func init() {
	var err error
	port, err = utils.MustGet("MUSIC_LEARNER_SERVER_PORT", "7777")
	if err != nil {
		log.Panicln(err)
	}
}

// Run web server

func Run() {
	utils.DefineLogger()

	router := gin.Default()
	db.OpenConnection()
	// Setup routes
	router.GET("/ping/:name", handlers.PingWithName())
	router.GET("/ping", handlers.Ping())
	router.GET("/song", handlers.SongHandler())
	router.GET("/mongo/save", handlers.GenerateAndSaveSongHandler())
	router.GET("/mongo/get", handlers.GetAnyMongo())
	router.GET("/note/random", handlers.RandomNoteHandler())

	log.Fatalln(router.Run(":" + port))
}
