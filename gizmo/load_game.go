package gizmo

import (
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
)

func LoadGameKeys(path string) (game *Game, err error) {
	gameKeys, _ := access.LoadGameKeys(filepath.Join(path, "game_keys.yaml"))
	var record = &model.Record{
		Major: gameKeys.Major,
		Minor: gameKeys.Minor,
	}
	return LoadGame(path, record)
}

func LoadGame(folder string, record *model.Record) (game *Game, err error) {
	game = &Game{}

	var path string
	path, err = filepath.Abs(folder)
	if err != nil {
		log.Println(err)
		return
	}

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
