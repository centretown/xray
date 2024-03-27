package tools

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	leftMargin   = 20
	rightMargin  = 20
	topMargin    = 50
	bottomMargin = 20
)

type Runner struct {
	Width  int32
	Height int32
	FPS    int32

	Actors []*Actor
}

func NewRunner(width int32, height int32, fps int32) *Runner {
	runr := &Runner{
		Height: height,
		Width:  width,
		FPS:    fps,
		Actors: make([]*Actor, 0),
	}

	return runr
}

func (runr *Runner) Add(d CanDraw, a CanMove, after float64) {
	runr.Actors = append(runr.Actors, NewActor(d, a, after))
}

func (runr *Runner) AddBouncingBalls() {
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

	viewPort := runr.GetViewPort()

	runr.Add(NewBall(60, colors[0]), NewBouncer(viewPort, 60, 60), 0)
	runr.Add(NewBall(40, colors[1]), NewBouncer(viewPort, 40, 40), 1)
	runr.Add(NewBall(30, colors[2]), NewBouncer(viewPort, 30, 30), 2)
	runr.Add(NewBall(20, colors[3]), NewBouncer(viewPort, 20, 20), 3)

}

func (runr *Runner) Refresh(current float64) {
	viewPort := runr.GetViewPort()
	for _, run := range runr.Actors {
		run.Resize(viewPort, current)
	}
}

func (runr *Runner) SetupWindow(title string) {
	rl.SetTraceLogLevel(rl.LogInfo)
	rl.InitWindow(runr.Width, runr.Height, title)
	rl.SetTargetFPS(runr.FPS)
	rl.SetWindowState(rl.FlagWindowResizable)
}

func (runr *Runner) GetMessageBox() rl.RectangleInt32 {
	v := runr.GetViewPort()
	v.Y = v.Height - 80
	v.Height -= v.Y
	return v
}

func (runr *Runner) GetViewPort() rl.RectangleInt32 {
	rw := rl.GetRenderWidth()
	if rw > 0 {
		return rl.RectangleInt32{
			X:      0,
			Y:      0,
			Width:  int32(rw),
			Height: int32(rl.GetRenderHeight()),
		}
	}

	return rl.RectangleInt32{
		X:      leftMargin,
		Y:      topMargin,
		Width:  runr.Width,
		Height: runr.Height,
	}
}
