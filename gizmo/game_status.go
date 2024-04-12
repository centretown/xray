package gizmo

import (
	"fmt"
	"image/color"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"gopkg.in/yaml.v3"
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
	rl.DrawText(text, mb.X, mb.Y+mb.Height-38, 21, yellow)

	if gs.Capturing {
		rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.CaptureCount),
			mb.X, mb.Y+mb.Height-70, 21, yellow)
	}
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.SetViewPortFromWindow()
	for _, run := range gs.actors {
		run.Refresh(current, viewPort)
	}
}

func (gs *Game) SetViewPortFromWindow() rl.RectangleInt32 {
	return gs.SetViewPort(int32(rl.GetRenderWidth()),
		int32(rl.GetRenderHeight()))
}

func (gs *Game) SetViewPort(rw, rh int32) rl.RectangleInt32 {
	gs.Width = rw
	gs.Height = rh
	return gs.GetViewPort()
}

func (gs *Game) GetViewPort() rl.RectangleInt32 {
	return rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  gs.Width,
		Height: gs.Height - msg_height,
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

func (game *Game) Dump() {
	buf, _ := yaml.Marshal(game)
	log.Println(string(buf))

	for _, mv := range game.Actors() {
		buf, _ = yaml.Marshal(mv)
		log.Println(string(buf))
		buf, _ = yaml.Marshal(mv.GetDrawer())
		log.Println(string(buf))
	}

	for _, dr := range game.Drawers() {
		buf, _ = yaml.Marshal(dr)
		log.Println(string(buf))
	}
}
