package gizmo

import (
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadGame(driver, path string, record *model.Record) *Game {
	gameData := dbg.NewGameData(driver, path)
	gameData.Open()
	if gameData.Err != nil {
		panic(gameData.Err)
	}
	defer gameData.Close()

	if gameData.Err != nil {
		panic(gameData.Err)
	}

	game := NewGame()
	game.Record = record

	gameData.Load(game)
	if gameData.Err != nil {
		panic(gameData.Err)
	}

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(game.Width, game.Height, game.Record.Title)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(game.FPS)

	game.SetColors()

	game.Refresh(rl.GetTime())

	return game
}
