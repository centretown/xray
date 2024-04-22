package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
)

type PaneItem struct {
	Rectangle rl.Rectangle
}

type Pane struct {
	model.RecorderClass[PaneItem]
}

func NewPaneFromRecord(record *model.Record) *Pane {
	pane := &Pane{}
	model.Decode(pane, record)
	return pane
}

func NewPane(rectangle rl.Rectangle) *Pane {
	pane := &Pane{}
	item := &pane.Content
	item.Rectangle = rectangle
	model.InitRecorder[Pane](pane,
		class.Pane.String(), int32(class.Pane))
	return pane
}
