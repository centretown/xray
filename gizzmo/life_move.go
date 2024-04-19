package gizzmo

import (
	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func NewLifeMoverFromRecord(record *model.Record) model.Recorder {
	lg := &GridMover[int8]{}
	model.Decode(lg, record)
	lg.init(true)
	lg.Refresh(rl.GetTime(), lg.Content.Bounds)
	return lg
}
