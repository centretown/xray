package main

import (
	"flag"
	"image/color"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/cmdl"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/gizmo"
)

var (
	installDir  = "/home/dave/xray/life_01/"
	installFile = installDir + "xray_game.db"
	installKeys = installDir + "game_keys.yaml"
)

func init() {
	cmdl.Path = ":memory:"
	cmdl.Setup()
}

func main() {
	flag.Parse()
	if cmdl.Path == "install" {
		cmdl.Path = installFile
	}
	createLife(cmdl.Path)
}

func createLife(fname string) {
	data := dbg.NewGameData("sqlite3", fname)
	data.Open()
	if data.Err != nil {
		panic(data.Err)
	}
	defer data.Close()

	if data.Err != nil {
		panic(data.Err)
	}

	data.Create()

	const (
		baseInterval = .02
		screenWidth  = 800
		screenHeight = 450
		fps          = 20
		captureFps   = 25
	)

	game := gizmo.NewGameSetup(screenWidth, screenHeight, fps)
	viewPort := game.GetViewPort()

	cells := gizmo.NewCells(viewPort.Width, viewPort.Height, 12)
	game.AddDrawer(cells)
	game.FrameRate = fps
	game.FixedSize = true
	cells.Colors = []color.RGBA{gizmo.White, gizmo.Blue, gizmo.Green}
	game.AddColors(cells.Colors)

	data.Save(game)
	if data.Err != nil {
		panic(data.Err)
	}

	access.SaveGameKeys(installKeys,
		access.NewGameKeys(game.Record.Major,
			game.Record.Minor))

	game.Dump()
}
