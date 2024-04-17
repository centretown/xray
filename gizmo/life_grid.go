package gizmo

import "github.com/centretown/xray/model"

func NewLifeGridFromRecord(record *model.Record) model.Recorder {
	lg := &NumberGrid[int8]{}
	model.Decode(lg, record)
	return lg
}
