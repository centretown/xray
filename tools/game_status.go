package tools

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	msg_height = 80
	min_width  = 200
	min_height = 280
)

func (gs *Game) DrawStatus() {
	mb := gs.GetMessageBox()
	rl.DrawLine(mb.X, mb.Y, mb.Width, mb.Y, color.RGBA{255, 0, 0, 255})

	monitor := rl.GetCurrentMonitor()

	text := fmt.Sprintf("FPS:%3d, Monitor:%1d (%4d/%4d %3d), View: %4dx%4d, Capture Count:%4d",
		rl.GetFPS(),
		monitor, rl.GetMonitorWidth(monitor),
		rl.GetMonitorHeight(monitor), rl.GetMonitorRefreshRate(monitor),
		rl.GetScreenWidth(), rl.GetScreenHeight(),
		gs.captureStart)

	yellow := color.RGBA{255, 255, 0, 255}
	rl.DrawText(text, mb.X, mb.Y+mb.Height-22, 16, yellow)

	if gs.Capturing {
		rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.CaptureCount),
			mb.X, mb.Y+32, 16, yellow)
	}
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.GetViewPort()
	for _, run := range gs.movers {
		run.Refresh(current, viewPort)
	}
}

func (gs *Game) GetViewPort() rl.RectangleInt32 {
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
		X:      0,
		Y:      0,
		Width:  min_width,
		Height: min_height - msg_height,
	}
}

func (gs *Game) GetMessageBox() (rect rl.RectangleInt32) {
	rw := int32(rl.GetRenderWidth())
	rh := int32(rl.GetRenderHeight())
	rect.X = 0
	rect.Width = rw
	rect.Y = rh - msg_height
	rect.Height = msg_height
	return
}

func (gs *Game) Dump() {
}
