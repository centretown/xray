package gizmo

import (
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var def_dir = "/home/dave/xray/"

func LoadGame(path string, record *model.Record) *Game {
	data := dbg.NewGameData("sqlite3", path+"xray_game.db")
	if data.Open().Err != nil {
		panic(data.Err)
	}
	defer data.Close()

	if data.Err != nil {
		panic(data.Err)
	}

	game := NewGame()
	game.Record = record
	game.path = path

	data.Load(game)
	if data.Err != nil {
		panic(data.Err)
	}

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(game.Width, game.Height, game.Record.Title)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(game.FPS)

	game.SetColors()
	game.Refresh(rl.GetTime())
	return game
}
