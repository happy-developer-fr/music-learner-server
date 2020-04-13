package db

import (
	"encoding/json"
	"github.com/happy-developer-fr/musical-notation/pkg/song"
	"github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
)

func (SongDb) TableName() string {
	return "song"
}

func From(song song.Song) (SongDb, error) {
	jsonNotes, err := json.Marshal(song.Notes)
	if err != nil {
		return SongDb{}, err
	}

	return SongDb{
		SongId: uuid.NewV4(),
		Notes:  postgres.Jsonb{RawMessage: jsonNotes},
	}, nil
}

type SongDb struct {
	SongId uuid.UUID      `gorm:"type:uuid;primary_key;Column:song_id"`
	Notes  postgres.Jsonb `gorm:"type:jsonb;Column:notes"`
}
