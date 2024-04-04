package main

import (
	"image/color"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/tools"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gp := &gpads.GPads{}
	gs := setup(gp)
	gs.Run()
}

func setup(gp pad.Pad) *tools.Game {

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(screenWidth, screenHeight, "2d")
	rl.SetWindowState(rl.FlagWindowResizable)

	gs := tools.NewGame(gp, fps)
	viewPort := gs.GetViewPort()

	hole := tools.NewPicture("polar.png").Load()
	bouncer := tools.NewMover(hole, viewPort, 10, 10, 10)
	gs.AddActor(bouncer, 6)

	ball := tools.NewCircle(20, tools.Cyan)
	bouncer = tools.NewMover(ball, viewPort, 200, 100, 0)
	gs.AddActor(bouncer, 1)

	head := tools.NewPicture("head_300.png").Load()
	bouncer = tools.NewMover(head, viewPort, 70, 140, 1.75)
	gs.AddActor(bouncer, 8)

	gander := tools.NewPicture("gander.png").Load()
	bouncer = tools.NewMover(gander, viewPort, 300, 300, 0.5)
	gs.AddActor(bouncer, 4)

	// generate palette and color map for paletted images
	pal, colorMap :=
		tools.CreatePaletteFromTextures(color.RGBA{}, fixedPalette, gs.Actors()...)
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
	tools.White,
	tools.Black,
	tools.Red,
	tools.Yellow,
	tools.Green,
	tools.Cyan,
	tools.Blue,
	tools.Magenta,
}
