package tools

import (
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
	Actors []*Actor
}

func NewRunner(width int32, height int32) *Runner {
	runr := &Runner{
		Height: height,
		Width:  width,
		Actors: make([]*Actor, 0),
	}

	return runr
}

func (runr *Runner) Add(d Drawable, a Moveable, after float64) {
	runr.Actors = append(runr.Actors, NewActor(d, a, after))
}

func (runr *Runner) Refresh(current float64) {
	viewPort := runr.GetViewPort()
	for _, run := range runr.Actors {
		run.Resize(viewPort, current)
	}
}

func (runr *Runner) SetupWindow(title string) {
	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(runr.Width, runr.Height, title)
	rl.SetWindowState(rl.FlagWindowResizable)
}

func (runr *Runner) GetMessageBox() (rect rl.RectangleInt32) {
	rw := int32(rl.GetRenderWidth())
	rh := int32(rl.GetRenderHeight())
	rect.X = 0
	rect.Width = rw
	rect.Y = rh - msg_height
	rect.Height = msg_height
	return
}

const (
	msg_height = 80
	min_width  = 200
	min_height = 280
)

func (runr *Runner) GetViewPort() rl.RectangleInt32 {
	rw := rl.GetRenderWidth()
	rh := rl.GetRenderHeight()

	if rw >= min_width && rh >= min_height {
		return rl.RectangleInt32{
			X:      0,
			Y:      0,
			Width:  int32(rw),
			Height: int32(rh - msg_height),
		}
	}

	return rl.RectangleInt32{
		X:      leftMargin,
		Y:      topMargin,
		Width:  min_width,
		Height: min_height - msg_height,
	}
}
