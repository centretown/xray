package gizmo

import (
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
)

func LoadGameKey(path string) (game *Game, err error) {
	gameKeys, _ := access.LoadGameKey(filepath.Join(path, "game_keys.yaml"))
	var record = &model.Record{
		Major: gameKeys.Major,
		Minor: gameKeys.Minor,
	}
	return LoadGame(path, record)
}

func LoadGame(folder string, record *model.Record) (game *Game, err error) {
	game = &Game{}
	path := filepath.Clean(folder)
	data := dbg.NewGameData("sqlite3", filepath.Join(path, "xray_game.db"))
	data.Open()

	defer func() {
		if data.Err != nil {
			err = data.Err
			log.Println(data.Err)
			return
		}
		data.Close()
	}()

	record = data.GetItem(record.Major, record.Minor)
	if data.HasErrors() {
		return
	}

	game.Setup(record, path)

	data.Load(game)
	if data.HasErrors() {
		return
	}

	return
}
