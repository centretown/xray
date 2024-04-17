package gizmo

import "github.com/centretown/xray/model"

func NewLifeMoverFromRecord(record *model.Record) model.Recorder {
	lg := &GridMover[int8]{}
	model.Decode(lg, record)
	return lg
}
