package db

import (
	"errors"
	"fmt"
	"github.com/happy-developer-fr/musical-notation/pkg/song"
)

func SaveSong(song song.Song) error {
	if DataBase.NewRecord(song) == false {
		errorString := fmt.Sprint("Song already exists")
		return errors.New(errorString)
	}
	return nil
}

func CountSong() int {
	count := 0
	DataBase.Count(&count)
	return count
}
