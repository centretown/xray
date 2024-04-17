package gizzmo

import (
	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewLifeMoverFromRecord(record *model.Record) model.Recorder {
	lg := &GridMover[int8]{}
	model.Decode(lg, record)
	lg.init(true)
	lg.Refresh(rl.GetTime(), rl.Vector4{
		X: lg.Content.Rectangle.Width,
		Y: lg.Content.Rectangle.Height,
	})
	return lg
}
