package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/happy-developer-fr/music-learner-server/internal/db"
	"github.com/happy-developer-fr/musical-notation/pkg/note"
	"github.com/happy-developer-fr/musical-notation/pkg/pitch"
	"github.com/happy-developer-fr/musical-notation/pkg/song"
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

func GenerateAndSaveSongHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db.SaveSong(randomSong(12))
	}
}

func GetAnyMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aSong song.Song
		db.DataBase.First(&aSong)
		c.JSON(http.StatusOK, aSong)
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
