package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/happy-developer-fr/musical-notation/pkg/note"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
	"github.com/happy-developer-fr/musical-notation/pkg/song"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func randomSong(n int) song.Song {
	var songNotes []note.Note
	for i := 0; i < n; i++ {
		songNotes = append(songNotes, note.RandomNote([]pitch.Pitch{pitch.C, pitch.G}))
	}
	return song.Song{Notes: songNotes}
}

func SongHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := c.Query("mode")
		randomSong := randomSong(12)
		if mode == "compact" {
			c.JSON(http.StatusOK, randomSong.Print())
		} else {
			c.JSON(http.StatusOK, randomSong)
		}
	}
}

func GenerateAndSaveSongHandler(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := collection.InsertOne(context.TODO(), randomSong(12)); err != nil {
			log.Fatalln("Error when inserting in mongo" + err.Error())
		}
	}
}

func GetAnyMongo(collection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		findOptions := options.Find()

		var aSong song.Song
		count, errCount := collection.CountDocuments(c, bson.D{{}})
		n := rand.Int63n(count)
		fmt.Println(n)
		findOptions.SetLimit(-1)
		findOptions.SetSkip(1)

		if errCount != nil {
			c.String(http.StatusInternalServerError, errCount.Error())
		} else {
			if err := collection.FindOne(c, bson.D{{}}).Decode(&aSong); err != nil {
				c.String(http.StatusInternalServerError, "error on find", err.Error())
			} else {
				c.JSON(http.StatusOK, aSong)
			}
		}
	}
}

func RandomNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := c.Query("mode")
		aNote := note.RandomNote([]pitch.Pitch{pitch.C, pitch.G})
		if mode == "compact" {
			c.JSON(http.StatusOK, aNote.Print())
		} else {
			c.JSON(http.StatusOK, aNote)
		}
	}
}

func SongNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		position := c.Param("pos")
		mode := c.Query("mode")
		pos, err := strconv.Atoi(position)
		if err != nil {
			c.String(http.StatusBadRequest, "pos is not an int")
		}
		randomSong := randomSong(12)
		if pos < 0 || pos > len(randomSong.Notes)-1 {
			c.String(http.StatusBadRequest, "pos is out of range")
		}
		if mode == "compact" {
			c.JSON(http.StatusOK, randomSong.Notes[pos].Print())
		} else {
			c.JSON(http.StatusOK, randomSong.Notes[pos])
		}
	}
}
