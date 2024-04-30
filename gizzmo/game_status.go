package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
	"gopkg.in/yaml.v3"
)

func (gs *Game) DrawStatus() {
	gs.RefreshEnvironment()
	var (
		content = &gs.Content
		row     = content.Layout.Fontsize
	)

	if content.commandState {
		row += gs.drawOutputs(row, content.Options)
	}

	if content.capturing {
		gs.drawOutputs(row, content.Capture)
	}
}

func (gs *Game) drawOutputs(row int32, notes *notes.Notebook) int32 {
	var layout = gs.Content.Layout
	return layout.Layout(row, notes, gs.Content.Language,
		func(y int32, label string, labelColor color.RGBA,
			value string, valueColor color.RGBA) {
			rl.DrawText(label, int32(layout.LabelX), int32(y), layout.Fontsize, labelColor)
			rl.DrawText(value, int32(layout.ValueX), int32(y), layout.Fontsize, valueColor)
		})
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
