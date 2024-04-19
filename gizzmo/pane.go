package gizzmo

import (
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
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
