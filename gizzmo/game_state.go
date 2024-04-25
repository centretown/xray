package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
	"gopkg.in/yaml.v3"
)

const (
	msg_font_size = 20
	msg_label_X   = msg_font_size + 3
	msg_value_X   = msg_label_X + msg_font_size*15
	msg_Y         = msg_font_size + 3
)

var (
	msg_color_label_input   = color.RGBA{255, 0, 255, 255}
	msg_color_value_input   = color.RGBA{255, 255, 255, 255}
	msg_color_label_data    = color.RGBA{128, 0, 128, 255}
	msg_color_value_data    = color.RGBA{128, 128, 128, 255}
	msg_color_label_display = color.RGBA{255, 255, 0, 255}
	msg_color_value_display = color.RGBA{0, 255, 255, 255}
	msg_color_label         = color.RGBA{128, 128, 0, 255}
	msg_color_value         = color.RGBA{0, 128, 128, 255}
)

func (gs *Game) DrawStatus() {
	gs.UpdateNotes()
	content := &gs.Content

	if !content.commandState && !content.capturing {
		return
	}

	content.notes.Fetch()

	row := int(msg_Y)
	row += gs.drawOutputs(row, content.notes)

	if content.capturing {
		content.captureNotes.Fetch()
		gs.drawOutputs(row, content.captureNotes)
	}
}

func (gs *Game) drawOutputs(row int, tkl *notes.Notes) int {
	var (
		content                  = &gs.Content
		color_label, color_value color.RGBA
	)

	draw := func(i int, label, value string) {
		if i == content.note {
			if content.notes.List[i].CanAct() {
				color_label = msg_color_label_input
				color_value = msg_color_value_input
			} else {
				color_label = msg_color_label_display
				color_value = msg_color_value_display
			}
		} else {
			if content.notes.List[i].CanAct() {
				color_label = msg_color_label_data
				color_value = msg_color_value_data
			} else {
				color_label = msg_color_label
				color_value = msg_color_value
			}
		}
		rl.DrawText(label, int32(msg_label_X), int32(row), msg_font_size, color_label)
		rl.DrawText(value, int32(msg_value_X), int32(row), msg_font_size, color_value)
	}

	for i := range tkl.Outputs {
		tkl.Draw(i, draw)
		row += msg_font_size * 2
	}
	return row
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
