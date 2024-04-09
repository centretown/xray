package gizmo

import (
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// var _ model.Linker = (*CellMoverItem)(nil)

type CellMoverItem struct {
	Bounds     rl.RectangleInt32
	PixelRateX float64
	drawer     *Cells
}

type CellsMover struct {
	CellMoverItem
	Record *model.Record
}

// var _ Actor = (*CellsMover)(nil)
// var _ model.Recorder = (*CellsMover)(nil)

// var _ model.Linker = (*CellsMover)(nil)

// func (cm *CellsMover) GetDrawer() Drawer        { return cm.drawer }
// func (cm *CellsMover) GetRecord() *model.Record { return cm.Record }
// func (cm *CellsMover) GetItem() any             { return &cm.CellMoverItem }

// func (cm *CellsMover) Link(recs ...*model.Record) {
// 	err := MakeLink(cm.AddDrawer, 1, 1, recs...)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (cs *CellsMover) Children() (rcds []model.Recorder) {
// 	rcds = make([]model.Recorder, 0)
// 	r, ok := cs.drawer.(model.Recorder)
// 	if ok {
// 		rcds = append(rcds, r)
// 	}
// 	return
// }

// func (cm *CellsMover) AddDrawer(dr *Cells) {
// 	cm.drawer = dr
// }

// func (cm *CellsMover) Children() (rcds []model.Recorder)

// func (cm *CellsMover) Act(can_move bool, current float64)
// func (cm *CellsMover) Refresh(now float64, rect rl.RectangleInt32)
