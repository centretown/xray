package main

import (
	"image"
	"image/color"
	"xray/b2"
	"xray/capture"
	"xray/gpads"
	"xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BLACK uint8 = iota
	RED
	YELLOW
	GREEN
	CYAN
	BLUE
	MAGENTA
	TRANSPARENT
)

var pal = color.Palette{
	rl.Black,
	rl.Red,
	rl.Yellow,
	rl.Green,
	color.RGBA{R: 0, G: 255, B: 255, A: 0},
	rl.Blue,
	rl.Magenta,
	color.Transparent,
}

func main() {
	runr := tools.NewRunner(640, 400, 60)
	viewPort := runr.GetViewPort()

	runr.Add(tools.NewBall(30, pal[RED]), tools.NewBouncer(viewPort, 30, 30), 0)
	runr.Add(tools.NewBall(20, pal[YELLOW]), tools.NewBouncer(viewPort, 20, 20), 1)
	runr.Add(tools.NewBall(15, pal[GREEN]), tools.NewBouncer(viewPort, 15, 15), 2)
	runr.Add(tools.NewBall(10, pal[BLUE]), tools.NewBouncer(viewPort, 10, 10), 3)
	Run2d(runr)
}

func Run2d(runr *tools.Runner) {
	runr.SetupWindow("2d")

	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime())
		can_move int32   = 0
		pads             = gpads.NewGPads()
	)

	runr.Refresh(current)

	for !rl.WindowShouldClose() {

		current = rl.GetTime()
		can_move = b2.ToInt32(current > previous+interval)
		previous = float64(can_move) * interval

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, run := range runr.Actors {
			run.Animate(can_move, current)
		}
		rl.EndDrawing()

		ProcessInput(pads, current, .02)
	}

	rl.CloseWindow()
}

var (
	next       float64
	capturing  bool
	frameCount int
	stopChan   = make(chan int)
	scrChan    = make(chan image.Image)
)

func ProcessInput(pads *gpads.GPads, current float64, interval float64) {
	pads.BeginPad()
	if current > next {
		next = current + interval
		if capturing {
			scrChan <- rl.LoadImageFromScreen().ToImage()
			frameCount--
			if frameCount < 0 {
				capturing = false
				stopChan <- 1
			}
			return
		}

		for i := range pads.GetStickCount() {
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft) {
				capturing = true
				frameCount = 360
				go capture.CaptureGIF(stopChan, scrChan, pal, interval)
				// go capture.CaptureGIF(stopChan, scrChan, colorMap, pal)
				return
			}
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight) {
				capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
				return
			}
		}
	}
}
