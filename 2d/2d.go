package main

import (
	"image/color"

	"github.com/centretown/xray/rayl"
	"github.com/centretown/xray/tools"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl := &rayl.RayL{}

	gs := setup(rl)
	gs.Run(rl)
}

func setup(rl rayl.RunLib) *tools.Game {

	rl.LogWarnings()
	rl.InitWindow(screenWidth, screenHeight, "2d")
	rl.SetWindowResizble()

	gs := tools.NewGame(fps, rl)
	viewPort := gs.GetViewPort()

	hole := tools.NewPicture("polar.png").Load()
	bouncer := tools.NewBouncer(hole, viewPort, 10, 10, 10)
	gs.AddActor(bouncer, 6)

	ball := tools.NewBall(20, color.White)
	bouncer = tools.NewBouncer(ball, viewPort, 200, 100, 0)
	gs.AddActor(bouncer, 1)

	head := tools.NewPicture("head_300.png").Load()
	bouncer = tools.NewBouncer(head, viewPort, 70, 140, 1.75)
	gs.AddActor(bouncer, 8)

	gander := tools.NewPicture("gander.png").Load()
	bouncer = tools.NewBouncer(gander, viewPort, 300, 300, 0.5)
	gs.AddActor(bouncer, 4)

	// generate palette and color map for paletted images
	pal, colorMap :=
		tools.CreatePaletteFromTextures(color.RGBA{}, fixedPalette, gs.Actors...)
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
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},               //RED
	color.RGBA{255, 255, 0, 255},             //Yellow,
	color.RGBA{0, 255, 0, 255},               //Green
	color.RGBA{R: 0, G: 255, B: 255, A: 255}, //cyan
	color.RGBA{R: 0, G: 0, B: 255, A: 255},   //blue
	color.RGBA{R: 255, G: 0, B: 255, A: 255}, //Magenta
}
