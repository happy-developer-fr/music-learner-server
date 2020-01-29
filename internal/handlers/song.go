package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/read-music-learner/musical-notation/musical_notation"
	"gitlab.com/read-music-learner/musical-notation/musical_notation/pitch"
	"gitlab.com/read-music-learner/musical-notation/musical_notation/song"
	"net/http"
	"strconv"
)

func randomSong(n int) song.Song {
	var songNotes []musical_notation.Note
	for i := 0; i < n; i++ {
		songNotes = append(songNotes, musical_notation.RandomNote([]pitch.Pitch{pitch.C, pitch.G}))
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

func RandomNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := c.Query("mode")
		note := musical_notation.RandomNote([]pitch.Pitch{pitch.C, pitch.G})
		if mode == "compact" {
			c.JSON(http.StatusOK, note.Print())
		} else {
			c.JSON(http.StatusOK, note)
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
