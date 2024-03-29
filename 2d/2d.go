package main

import (
	"image/color"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WHITE uint8 = iota
	BLACK
	RED
	YELLOW
	GREEN
	CYAN
	BLUE
	MAGENTA
)

const (
	baseInterval = .02
	screenWidth  = 1280
	screenHeight = 720
	fps          = 30
	captureFps   = 25
)

// frames per second sent

// delay in 100ths of a second
// if delay = 4/100 s = 25 fps
///           1/100 s = 100 fps
//			  2=50

func main() {

	var fixedPal = color.Palette{
		rl.White,
		rl.Black,
		rl.Red,
		rl.Yellow,
		rl.Green,
		color.RGBA{R: 0, G: 255, B: 255, A: 255},
		rl.Blue,
		rl.Magenta,
	}

	runr := tools.NewRunner(screenWidth, screenHeight)

	runr.SetupWindow("2d")
	viewPort := runr.GetViewPort()

	gs := NewGameState(fps)

	hole := tools.NewPicture(rl.LoadTexture("hole.png"), -1)
	gs.actors = append(gs.actors, hole)
	bouncer := tools.NewBouncer(hole.Rect(), viewPort, 0, 0)
	runr.Add(gs.actors[0], bouncer, 6)

	ball := tools.NewBall(40, rl.Green)
	bouncer = tools.NewBouncer(ball.Rect(), viewPort, 50, 200)
	runr.Add(ball, bouncer, 1)

	head := tools.NewPicture(rl.LoadTexture("head_90.png"), 10)
	gs.actors = append(gs.actors, head)
	bouncer = tools.NewBouncer(head.Rect(), viewPort, 240, 240)
	runr.Add(gs.actors[1], bouncer, 8)

	gander := tools.NewPicture(rl.LoadTexture("gander.png"), 2)
	gs.actors = append(gs.actors, gander)
	bouncer = tools.NewBouncer(gander.Rect(), viewPort, 66, 66)
	runr.Add(gs.actors[2], bouncer, 4)

	// generate palette and color map for paletted images
	gs.pal, gs.colorMap = createPaletteFromTextures(fixedPal, gs.actors...)
	// time.Sleep(time.Second * 2)

	rl.SetTargetFPS(gs.fps)
	runr.Refresh(rl.GetTime())
	gs.Dump()

	defer func() {
		for i := range gs.actors {
			rl.UnloadTexture(gs.actors[i].Texture2D)
		}
		rl.CloseWindow()
	}()

	for !rl.WindowShouldClose() {

		gs.current = rl.GetTime()

		if rl.IsWindowResized() {
			runr.Refresh(gs.current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for _, run := range runr.Actors {
			run.Animate(!gs.paused, gs.current)
		}

		gs.DrawStatus(runr)

		rl.EndDrawing()

		gs.ProcessInput()
	}
}

func createPaletteFromTextures(pal color.Palette, heads ...*tools.Picture) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	x := int32(0)
	for _, head := range heads {
		rl.DrawTexture(head.Texture2D, x, 0, rl.White)
		x += head.Rect().Width + 120
	}
	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(pal, pic, 256)
}
