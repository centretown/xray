package gizmo

import (
	"github.com/centretown/xray/gizmodb/model"
)

func NewLifeGridFromRecord(record *model.Record) model.Recorder {
	lg := &NumberGrid[int8]{}
	model.Decode(lg, record)

	// fmt.Println("SETUP CELLS")
	// lg.SetupCells()
	return lg
}
