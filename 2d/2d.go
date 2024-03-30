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

	run := tools.NewRunner(screenWidth, screenHeight)

	run.SetupWindow("2d")
	viewPort := run.GetViewPort()

	gs := NewGameState(fps)

	hole := tools.NewPicture(rl.LoadTexture("polar.png"))
	gs.textures = append(gs.textures, hole)
	bouncer := tools.NewBouncer(hole.Rect(), viewPort, 10, 10, 10)
	run.Add(gs.textures[0], bouncer, 6)

	ball := tools.NewBall(20, rl.Green)
	bouncer = tools.NewBouncer(ball.Rect(), viewPort, 200, 100, 0)
	run.Add(ball, bouncer, 1)

	head := tools.NewPicture(rl.LoadTexture("head_90.png"))
	gs.textures = append(gs.textures, head)
	bouncer = tools.NewBouncer(head.Rect(), viewPort, 240, 240, 2.75)
	run.Add(gs.textures[1], bouncer, 8)

	gander := tools.NewPicture(rl.LoadTexture("gander.png"))
	gs.textures = append(gs.textures, gander)
	bouncer = tools.NewBouncer(gander.Rect(), viewPort, 300, 300, 0.5)
	run.Add(gs.textures[2], bouncer, 4)

	// generate palette and color map for paletted images
	gs.pal, gs.colorMap = createPaletteFromTextures(fixedPal, gs.textures...)
	// time.Sleep(time.Second * 2)

	rl.SetTargetFPS(gs.fps)
	run.Refresh(rl.GetTime())
	gs.Dump()

	defer func() {
		for i := range gs.textures {
			rl.UnloadTexture(gs.textures[i].Texture2D)
		}
		rl.CloseWindow()
	}()

	loop(run, gs)
}

var BG = color.RGBA{R: 15, G: 0, B: 0, A: 255}

func loop(run *tools.Runner, gs *Game) {

	for !rl.WindowShouldClose() {

		gs.current = rl.GetTime()

		if rl.IsWindowResized() {
			run.Refresh(gs.current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(BG)

		for _, run := range run.Actors {
			run.Animate(!gs.paused, gs.current)
		}

		gs.DrawStatus(run)

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
