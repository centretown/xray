package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmodb/model"
)

func NewLifeMoverFromRecord(record *model.Record) model.Recorder {
	lg := &GridMover[int8]{}
	model.Decode(lg, record)
	lg.init(true)
	lg.Refresh(rl.GetTime(), lg.Content.Bounds)
	return lg
}
