package main

import (
	"image/color"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gizmo"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	gp := &gpads.GPads{}
	gs := setup(gp)
	gs.Run()
}

func setup(gp pad.Pad) *gizmo.Game {
	picd := "/home/dave/xray/test/pic/"

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(screenWidth, screenHeight, "2d")
	rl.SetWindowState(rl.FlagWindowResizable)

	gs := gizmo.NewGameSetup(screenWidth, screenHeight, fps)
	gs.SetPad(gp)
	viewPort := gs.SetViewPortFromWindow()

	hole := gizmo.NewTexture(picd + "polar.png").
		Load()
	bouncer := gizmo.NewMover(viewPort, 10, 10, 10).
		AddDrawer(hole)
	gs.AddMover(bouncer, 6)

	ball := gizmo.NewCircle(20, gizmo.Cyan)
	bouncer = gizmo.NewMover(viewPort, 200, 100, 0).
		AddDrawer(ball)
	gs.AddMover(bouncer, 1)

	head := gizmo.NewTexture(picd + "head_300.png").
		Load()
	bouncer = gizmo.NewMover(viewPort, 70, 140, 1.75).
		AddDrawer(head)
	gs.AddMover(bouncer, 8)

	gander := gizmo.NewTexture(picd + "gander.png").
		Load()
	bouncer = gizmo.NewMover(viewPort, 300, 300, 0.5).
		AddDrawer(gander)
	gs.AddMover(bouncer, 4)

	// generate palette and color map for paletted images
	pal, colorMap := gizmo.CreatePaletteFromTextures(color.RGBA{}, fixedPalette, gs)
	gs.SetColorPalette(color.RGBA{R: 0, G: 0, B: 0, A: 255}, pal, colorMap)

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
	gizmo.White,
	gizmo.Black,
	gizmo.Red,
	gizmo.Yellow,
	gizmo.Green,
	gizmo.Cyan,
	gizmo.Blue,
	gizmo.Magenta,
}
