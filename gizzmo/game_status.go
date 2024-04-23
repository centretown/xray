package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	msg "github.com/centretown/xray/message"
	"gopkg.in/yaml.v3"
)

const (
	msg_font_size = 21
	msg_X         = msg_font_size
	msg_Y         = msg_font_size
	msg_capture_Y = msg_Y + 2*msg_font_size
)

var (
	msg_color = color.RGBA{255, 255, 0, 255}
)

func (gs *Game) DrawStatus() {
	content := &gs.Content
	options := &msg.Options{Sep: ":", TokenSep: " "}

	monitor := rl.GetCurrentMonitor()

	text := msg.Message(
		options,
		&msg.Token{Item: msg.FPS, Format: "%d", Values: []any{rl.GetFPS()}},
		&msg.Token{Item: msg.Monitor, Format: "%d %dx%d %d%s",
			Values: []any{monitor,
				rl.GetMonitorWidth(monitor), rl.GetMonitorHeight(monitor),
				rl.GetMonitorRefreshRate(monitor), msg.Mhz}},
		&msg.Token{Item: msg.View, Format: "%dx%d",
			Values: []any{rl.GetScreenWidth(), rl.GetScreenHeight()}},
		&msg.Token{Item: msg.Duration, Format: "%q",
			Values: []any{content.CaptureDuration}},
		&msg.Token{Item: msg.Frames, Format: "%d",
			Values: []any{content.captureFrames}},
	)

	rl.DrawText(text, msg_X, msg_Y, msg_font_size, msg_color)

	if content.capturing {
		// text = fmt.Sprintf("Capturing... %4d", content.captureCount)
		text = msg.Message(options, &msg.Token{
			Item: msg.Capturing, Format: "...%d",
			Values: []any{content.captureCount}})

		rl.DrawText(text, msg_X, msg_capture_Y, msg_font_size, msg_color)
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
		Height: gs.Content.Height,
	}
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
