package gizmo

import (
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

func LoadGame(path string, record *model.Record) (game *Game, err error) {
	data := dbg.NewGameData("sqlite3", filepath.Join(path, "xray_game.db"))
	if data.Open().Err != nil {
		err = data.Err
		return
	}
	defer data.Close()

	record = data.GetItem(record.Major, record.Minor)
	if data.HasErrors() {
		err = data.Err
		return
	}

	game = &Game{}
	game.Setup(record, path)
	data.Load(game)
	if data.HasErrors() {
		err = data.Err
		return
	}

	return
}
