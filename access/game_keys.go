package access

import (
	"log"

	"github.com/centretown/xray/gizmodb/model"
	"github.com/google/uuid"
)

type GameKey struct {
	Major int64
	Minor int64
	ID    uuid.UUID
}

func NewGameKeys(major, minor int64) *GameKey {
	return &GameKey{
		Major: major,
		Minor: minor,
		ID:    model.RecordUUID(major, minor),
	}
}

type GameKeys []*GameKey

func LoadGameKey(path string) (key *GameKey, err error) {
	gks, err := LoadGameKeys(path)
	if err != nil || len(gks) == 0 {
		log.Fatal("count:", len(gks), err)
	}

	key = &GameKey{}
	*key = *gks[0]
	return
}

func LoadGameKeys(path string) (gks GameKeys, err error) {
	gks = make(GameKeys, 0)
	err = Load(path, &gks)
	return
}

func SaveGameKeys(path string, gks GameKeys) error {
	return Save(path, &gks)
}

func SaveGameKey(path string, gk *GameKey) (err error) {
	var (
		oldkeys = make(GameKeys, 0)
		keys    = make(GameKeys, 0)
	)

	keys = append(keys, gk)
	oldkeys, err = LoadGameKeys(path)
	if err == nil {
		keys = append(keys, oldkeys...)
	}

	err = SaveGameKeys(path, keys)
	return
}
