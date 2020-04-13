package db

import (
	"github.com/happy-developer-fr/musical-notation/pkg/song"
	uuid "github.com/satori/go.uuid"
)

func SaveSong(song song.Song) (uuid.UUID, error) {
	songDb, err := From(song)
	if err != nil {
		return uuid.Nil, err
	}
	DataBase.Create(&songDb)
	return songDb.SongId, nil
}

func CountSong() int {
	count := 0
	DataBase.Model(&SongDb{}).Count(&count)
	return count
}
