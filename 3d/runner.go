package main

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmo"
)

const (
	leftMargin   = 20
	rightMargin  = 20
	topMargin    = 50
	bottomMargin = 20
)

type Run3D struct {
	Width  int32
	Height int32
	Actors []gizzmo.Mover
}

func NewRunner(width int32, height int32) *Run3D {
	runr := &Run3D{
		Height: height,
		Width:  width,
		Actors: make([]gizzmo.Mover, 0),
	}

	return runr
}

func (runr *Run3D) Add(a gizzmo.Mover, after float64) {
	runr.Actors = append(runr.Actors, a)
}

func (runr *Run3D) Refresh(current float64) {

	viewPort := rl.RectangleInt32{X: 0, Y: 0, Width: int32(rl.GetRenderWidth()),
		Height: int32(rl.GetRenderHeight())}
	vp := viewPort.ToFloat32()

	for _, mover := range runr.Actors {
		mover.Refresh(current, rl.Vector4{X: vp.X, Y: vp.Y})
	}
}

func (runr *Run3D) SetupWindow(title string) {
	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(runr.Width, runr.Height, title)
	rl.SetWindowState(rl.FlagWindowResizable)
}

func (runr *Run3D) GetMessageBox() (rect rl.RectangleInt32) {
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

func (runr *Run3D) GetViewPort() rl.RectangleInt32 {
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
