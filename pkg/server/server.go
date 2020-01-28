package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/read-music-learner/music-learner-server/internal/handlers"
	"gitlab.com/read-music-learner/music-learner-server/pkg/utils"
	"io"
	"log"
	"os"
)

var (
	host string
	port string
)

const mode = os.ModePerm

func init() {
	var err error
	host, err = utils.MustGet("MUSIC_LEARNER_SERVER_HOST", "localhost")
	if err != nil {
		log.Panicln(err)
	}
	port, err = utils.MustGet("MUSIC_LEARNER_SERVER_PORT", "7777")
	if err != nil {
		log.Panicln(err)
	}
}

// Run web server

func Run() {
	var f *os.File
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		_ = os.Mkdir("log", mode)
	}

	f, _ = os.Create("log/music-learner.log")

	var offset int64 = 0
	if _, seekErr := f.Seek(offset, io.SeekEnd); seekErr != nil {
		log.Panicln("unable to go to end of file music-learner.log", seekErr)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, f))

	router := gin.Default()

	// Setup routes
	router.GET("/ping/:name", handlers.PingWithName())
	router.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(router.Run(host + ":" + port))
}
