package server

import (
	"github.com/gin-gonic/gin"
	"github.com/happy-developer-fr/music-learner-server/internal"
	"github.com/happy-developer-fr/music-learner-server/internal/handlers"
	"github.com/happy-developer-fr/music-learner-server/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
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
	var client *mongo.Client
	if mongoUrl, err := utils.MustGet("MONGO_URL", "mongodb://localhost:27017"); err != nil {
	} else {
		client = internal.ConnectMongo(mongoUrl)
	}
	collection := client.Database("musical_notation").Collection("songs")
	router := gin.Default()

	// Setup routes
	router.GET("/ping/:name", handlers.PingWithName())
	router.GET("/ping", handlers.Ping())
	router.GET("/song", handlers.SongHandler())
	router.GET("/mongo/save", handlers.GenerateAndSaveSongHandler(collection))
	router.GET("/mongo/get", handlers.GetAnyMongo(collection))
	router.GET("/note/random", handlers.RandomNoteHandler())

	log.Fatalln(router.Run(":" + port))
}
