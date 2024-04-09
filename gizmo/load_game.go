package gizmo

import (
	"github.com/centretown/xray/access"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadGameKeys(path string) (game *Game, err error) {
	gameKeys, _ := access.LoadGameKeys(path + "game_keys.yaml")
	var record = &model.Record{
		Major: gameKeys.Major,
		Minor: gameKeys.Minor,
	}
	return LoadGame(path, record)
}

func LoadGame(path string, record *model.Record) (game *Game, err error) {
	data := dbg.NewGameData("sqlite3", path+"xray_game.db")
	if data.Open().Err != nil {
		err = data.Err
		return
	}
	defer data.Close()

	record = data.GetItem(record.Major, record.Minor)
	if data.Err != nil {
		err = data.Err
		return
	}

	game = &Game{}
	game.Setup(record, path)
	data.Load(game)
	if data.Err != nil {
		err = data.Err
		return
	}

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(game.Width, game.Height, game.Record.Title)
	if !game.FixedSize {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
	rl.SetTargetFPS(game.FrameRate)
	// fmt.Println("game.FrameRate", game.FrameRate)

	game.SetColors()
	game.Refresh(rl.GetTime())
	return
}
