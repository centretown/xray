package main

import (
	"fmt"
	"github/centretown/xray/b2"
	"github/centretown/xray/capture"
	"github/centretown/xray/tools"
	"image/color"

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
)

// var (
// 	next       float64
// 	capturing  bool
// 	paused     bool
// 	frameCount int
// 	stopChan   = make(chan int)
// 	scrChan    = make(chan image.Image)
// )

func main() {

	var fixedPal = color.Palette{
		rl.White,
		rl.Black,
		rl.Red,
		rl.Yellow,
		rl.Green,
		color.RGBA{R: 0, G: 255, B: 255, A: 0},
		rl.Blue,
		rl.Magenta,
	}

	runr := tools.NewRunner(screenWidth, screenHeight, fps)
	runr.SetupWindow("2d")
	viewPort := runr.GetViewPort()

	gs := NewGameState()

	AddBouncingBall(runr, 40, rl.Red, 0)

	head := rl.LoadTexture("head_90.png")
	gs.actors = append(gs.actors, tools.NewTextureDrawer(head))
	bouncer := tools.NewBouncer(viewPort, head.Width, head.Height)
	runr.Add(gs.actors[0], bouncer, 8)

	gander := rl.LoadTexture("gander.png")
	gs.actors = append(gs.actors, tools.NewTextureDrawer(gander))
	bouncer = tools.NewBouncer(viewPort, gander.Width, gander.Height)
	runr.Add(gs.actors[1], bouncer, 4)

	runr.Refresh(gs.current)

	gs.pal, gs.colorMap = createPaletteFromTextures(fixedPal, gs.actors...)

	for !rl.WindowShouldClose() {

		gs.ProcessInput()

		gs.current = rl.GetTime()

		bMove := gs.current >= gs.previous+gs.interval
		gs.can_move = b2.To[int32](bMove && !gs.paused)
		moveFloat := float64(gs.can_move)
		gs.previous = moveFloat*gs.interval + moveFloat*gs.current

		if rl.IsWindowResized() {
			runr.Refresh(gs.current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for _, run := range runr.Actors {
			run.Animate(gs.can_move, gs.current)
		}

		mb := runr.GetMessageBox()
		gs.DrawStatus(runr)

		if gs.capturing && bMove {
			rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.captureCount),
				mb.X, mb.Y+32, 20, rl.Green)
		}
		rl.EndDrawing()

		if gs.capturing && bMove {
			gs.GIFCapture()
		}
	}

	rl.UnloadTexture(head)
	rl.CloseWindow()
}

func createPaletteFromTextures(pal color.Palette, heads ...*tools.TextureDrawer) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	x := int32(0)
	for _, head := range heads {
		rl.DrawTexture(head.Texture2D, x, 0, rl.White)
		// h.Draw(x, 0)
		x += head.Width() + 120
	}
	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(pal, pic)
}

func AddBouncingBall(runr *tools.Runner, radius int32, clr color.RGBA, layer float64) {
	runr.Add(tools.NewBall(radius, clr), tools.NewBouncer(runr.GetViewPort(), radius, radius), layer)
}
