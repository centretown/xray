package model

import (
	"encoding/json"
	"log"
)

type Recorder interface {
	GetRecord() *Record
	GetItem() any
	// Decode() error
	// Make(*Record) Recorder
}

type Parent interface {
	Recorder
	LinkChild(Recorder)
	Children() []Recorder
}

type RecorderClass[T any] struct {
	Record  Record
	Content T
}

func SetupRecorder[T any](gd Recorder, name string, classn int32) {
	SetupRecord(gd.GetRecord(), name, classn, gd.GetItem(), JSON)
}

func (rr *RecorderClass[T]) GetRecord() *Record {
	return &rr.Record
}

func (rr *RecorderClass[T]) GetItem() any {
	return &rr.Content
}

func Decode(recorder Recorder, rec *Record) (err error) {
	record := recorder.GetRecord()
	record.Copy(rec)
	err = json.Unmarshal([]byte(record.Content), recorder.GetItem())
	if err != nil {
		log.Println(record.Content)
		panic(err)
	}
	return
}
