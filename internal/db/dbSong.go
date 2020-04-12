package db

import uuid "github.com/satori/go.uuid"

func (SongDb) TableName() string {
	return "song"
}

type SongDb struct {
	SongId uuid.UUID `gorm:"type:uuid;primary_key;Column:song_id"`
}
