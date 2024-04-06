package main

import (
	"image/color"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/game"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gp := &gpads.GPads{}
	gs := setup(gp)
	gs.Run()
}

func setup(gp pad.Pad) *game.Game {
	picd := "/home/dave/xray/test/pic/"

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(screenWidth, screenHeight, "2d")
	rl.SetWindowState(rl.FlagWindowResizable)

	gs := game.NewGame(gp, screenWidth, screenHeight, fps)
	viewPort := gs.SetViewPortFromWindow()

	hole := game.NewTexture(picd + "polar.png").
		Load()
	bouncer := game.NewMover(viewPort, 10, 10, 10).
		AddDrawer(hole)
	gs.AddMover(bouncer, 6)

	ball := game.NewCircle(20, game.Cyan)
	bouncer = game.NewMover(viewPort, 200, 100, 0).
		AddDrawer(ball)
	gs.AddMover(bouncer, 1)

	head := game.NewTexture(picd + "head_300.png").
		Load()
	bouncer = game.NewMover(viewPort, 70, 140, 1.75).
		AddDrawer(head)
	gs.AddMover(bouncer, 8)

	gander := game.NewTexture(picd + "gander.png").
		Load()
	bouncer = game.NewMover(viewPort, 300, 300, 0.5).
		AddDrawer(gander)
	gs.AddMover(bouncer, 4)

	// generate palette and color map for paletted images
	pal, colorMap :=
		game.CreatePaletteFromTextures(color.RGBA{}, fixedPalette, gs.Movers()...)
	gs.SetColors(color.RGBA{}, pal, colorMap)

	rl.SetTargetFPS(gs.FPS)
	gs.Refresh(rl.GetTime())
	gs.Dump()

	return gs
}

const (
	baseInterval = .02
	screenWidth  = 1280
	screenHeight = 720
	fps          = 30
	captureFps   = 25
)

var BG = color.Black

const (
	TRANSPARENT = iota
	WHITE
	BLACK
	RED
	YELLOW
	GREEN
	CYAN
	BLUE
	MAGENTA
)

var fixedPalette = color.Palette{
	color.Transparent,
	game.White,
	game.Black,
	game.Red,
	game.Yellow,
	game.Green,
	game.Cyan,
	game.Blue,
	game.Magenta,
}
