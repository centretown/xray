package gizzmo

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
	item := &gs.Content
	mb := gs.GetMessageBox()
	rl.DrawLine(mb.X, mb.Y, mb.Width, mb.Y, color.RGBA{255, 0, 0, 255})

	monitor := rl.GetCurrentMonitor()

	text := fmt.Sprintf("FPS:%3d, Monitor:%1d (%4d/%4d %3d), View: %4dx%4d, Capture Count:%4d",
		rl.GetFPS(),
		monitor, rl.GetMonitorWidth(monitor),
		rl.GetMonitorHeight(monitor), rl.GetMonitorRefreshRate(monitor),
		rl.GetScreenWidth(), rl.GetScreenHeight(),
		item.CaptureStart)

	yellow := color.RGBA{255, 255, 0, 255}
	rl.DrawText(text, mb.X, mb.Y+mb.Height-38, 21, yellow)

	if item.Capturing {
		rl.DrawText(fmt.Sprintf("Capturing... %4d", item.CaptureCount),
			mb.X, mb.Y+mb.Height-70, 21, yellow)
	}
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.SetViewPort(float32(rl.GetRenderWidth()),
		float32(rl.GetRenderHeight()))

	for _, mover := range gs.Content.movers {
		mover.Refresh(current, rl.Vector4{
			X: viewPort.Width,
			Y: viewPort.Height})
	}
	for _, drawer := range gs.Content.drawers {
		drawer.Refresh(current, rl.Vector4{
			X: float32(viewPort.Width),
			Y: float32(viewPort.Height)})
	}
}

func (gs *Game) SetViewPort(rw, rh float32) rl.Rectangle {
	gs.Content.Width = rw
	gs.Content.Height = rh
	return gs.GetViewPort()
}

func (gs *Game) GetViewPort() rl.Rectangle {
	return rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  gs.Content.Width,
		Height: gs.Content.Height - msg_height,
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
