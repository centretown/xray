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
	installDir  = "/home/dave/xray/game_01/"
	installFile = installDir + "xray_game.db"
	installKeys = installDir + "game_keys.yaml"
	picd        = installDir + ""
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
	createGander(cmdl.Path)
}

func createGander(fname string) {
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

	var fixedPalette = []color.RGBA{
		gizmo.White,
		gizmo.Black,
		gizmo.Red,
		gizmo.Yellow,
		gizmo.Green,
		gizmo.Cyan,
		gizmo.Blue,
		gizmo.Magenta,
	}

	const (
		baseInterval = .02
		screenWidth  = 1280
		screenHeight = 720
		fps          = 30
		captureFps   = 25
	)

	game := gizmo.NewGameSetup(screenWidth, screenHeight, fps)
	viewPort := game.GetViewPort()

	hole := gizmo.NewTexture(picd + "polar.png")
	hole_mv := gizmo.NewMover(viewPort, 10, 10, 10)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, 6)

	ball := gizmo.NewCircle(20, gizmo.Cyan)
	ball_mv := gizmo.NewMover(viewPort, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, 1)

	head := gizmo.NewTexture(picd + "head_300.png")
	head_mv := gizmo.NewMover(viewPort, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, 8)

	gander := gizmo.NewTexture(picd + "gander.png")
	gander_mv := gizmo.NewMover(viewPort, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, 4)

	game.FixedPalette = fixedPalette
	data.Save(game)
	if data.Err != nil {
		panic(data.Err)
	}
	access.SaveGameKeys(installKeys,
		access.NewGameKeys(game.Record.Major,
			game.Record.Minor))

	game.Dump()
}
