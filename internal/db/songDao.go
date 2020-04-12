package db

import (
	"github.com/happy-developer-fr/musical-notation/pkg/song"
	uuid "github.com/satori/go.uuid"
)

func SaveSong(song song.Song) uuid.UUID {
	songDb := SongDb{SongId: uuid.NewV4()}
	DataBase.Create(&songDb)
	return songDb.SongId
}

func CountSong() int {
	count := 0
	DataBase.Count(&count)
	return count
}
