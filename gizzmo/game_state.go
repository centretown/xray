package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	msg "github.com/centretown/xray/message"
	"gopkg.in/yaml.v3"
)

const (
	msg_font_size = 20
	msg_label_X   = msg_font_size + 3
	msg_value_X   = msg_label_X + msg_font_size*10
	msg_Y         = msg_font_size + 3
)

var (
	msg_color_label = color.RGBA{255, 255, 0, 255}
	msg_color_value = color.RGBA{0, 255, 255, 255}
	msg_options     = &msg.Options{Sep: ":", TokenSep: " "}
)

func (gs *Game) DrawStatus() {
	content := &gs.Content
	if !content.commandState && !content.capturing {
		return
	}

	monitor := rl.GetCurrentMonitor()
	outputs := msg.Build(msg_options,
		&msg.Token{Label: msg.Monitor, Format: "%d %dx%d %d%s",
			Values: []any{monitor,
				rl.GetMonitorWidth(monitor), rl.GetMonitorHeight(monitor),
				rl.GetMonitorRefreshRate(monitor), msg.Mhz}},
		&msg.Token{Label: msg.View, Format: "%dx%d",
			Values: []any{rl.GetScreenWidth(), rl.GetScreenHeight()}},
		&msg.Token{Label: msg.Duration, Format: "%.3f",
			Values: []any{content.CaptureDuration}},
		&msg.Token{Label: msg.FrameRate, Format: "%d", Values: []any{rl.GetFPS()}},
	)

	y := int(msg_Y)
	y += drawOutputs(y, outputs)

	if content.capturing {
		outputs = msg.Build(msg_options,
			&msg.Token{Label: msg.Capturing, Format: "%d ... %d",
				Values: []any{content.captureTotal, content.captureCount}},
		)
		drawOutputs(y, outputs)
	}
}

func drawOutputs(y int, outputs []*msg.Output) int {
	for _, output := range outputs {
		rl.DrawText(output.Label, int32(msg_label_X), int32(y), msg_font_size, msg_color_label)
		rl.DrawText(output.Value, int32(msg_value_X), int32(y), msg_font_size, msg_color_value)
		y += msg_font_size * 2
	}
	return y
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
