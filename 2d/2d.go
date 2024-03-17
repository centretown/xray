package main

import (
	"image/color"
	"xray/b2"
	"xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var colors = []color.RGBA{
	rl.White,
	rl.Blue,
	rl.Yellow,
	rl.Red,
	rl.White,
	rl.Red,
	rl.White,
	rl.Yellow,
	rl.Lime,
	rl.DarkGreen,
}

func main() {
	runr := tools.NewRunner(1280, 720, 60)
	viewPort := runr.GetViewPort()

	runr.Add(tools.NewBall(60, colors), tools.NewBouncer(viewPort, 60, 60), 0)
	runr.Add(tools.NewBall(40, colors[6:]), tools.NewBouncer(viewPort, 40, 40), 1)
	runr.Add(tools.NewBall(30, colors[2:]), tools.NewBouncer(viewPort, 30, 30), 2)
	runr.Add(tools.NewBall(20, colors[4:]), tools.NewBouncer(viewPort, 20, 20), 3)
	Run2d(runr)
}

func Run2d(runr *tools.Runner) {
	runr.SetupWindow("2d")

	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime())
		can_move int32   = 0
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
	}

	rl.CloseWindow()
}
