package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
	"gopkg.in/yaml.v3"
)

func (gs *Game) DrawStatus() {
	gs.UpdateNotes()
	content := &gs.Content

	if !content.commandState && !content.capturing {
		return
	}

	row := content.layout.FontSize
	row += gs.drawOutputs(row, content.notes)

	if content.capturing {
		gs.drawOutputs(row, content.captureNotes)
	}
}

func (gs *Game) drawOutputs(row int32, notes *notes.Notes) int32 {

	var (
		layout = gs.Content.layout
	)

	layout.Layout(row, notes, gs.Content.language,
		func(y int32, label string, labelColor color.RGBA,
			value string, valueColor color.RGBA) {
			rl.DrawText(label, int32(layout.LabelX), int32(y), layout.FontSize, labelColor)
			rl.DrawText(value, int32(layout.ValueX), int32(y), layout.FontSize, valueColor)
		})
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
