package access

import (
	"github.com/centretown/xray/model"
	"github.com/google/uuid"
)

type GameKeys struct {
	Major int64
	Minor int64
	ID    uuid.UUID
}

func NewGameKeys(major, minor int64) *GameKeys {
	return &GameKeys{
		Major: major,
		Minor: minor,
		ID:    model.RecordUUID(major, minor),
	}
}

func LoadGameKeys(path string) (gk *GameKeys, err error) {
	return Load(path, &GameKeys{})
}

func SaveGameKeys(path string, gk *GameKeys) (err error) {
	return Save(path, gk)
}
