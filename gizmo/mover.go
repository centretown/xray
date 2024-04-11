package gizmo

import (
	"log"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Actor = (*Mover)(nil)
var _ model.Linker = (*Mover)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type MoverItem struct {
	Source       rl.RectangleInt32
	Bounds       rl.RectangleInt32
	PixelRateY   float64 // pixels per second
	PixelRateX   float64
	Rotation     float32
	RotationRate float32
	Axes         []*Axis
	drawer       Drawer
}

type Mover struct {
	MoverItem
	Record *model.Record
}

func NewMover(bounds rl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotationRate float32) *Mover {

	mv := &Mover{}
	mv.Bounds = bounds
	mv.PixelRateX = pixelRateX
	mv.PixelRateY = pixelRateY
	mv.Rotation = 0
	mv.RotationRate = rotationRate
	mv.Axes = make([]*Axis, 2)

	now := rl.GetTime()
	mv.Axes[0] = NewAxis(mv.Bounds.Width)
	mv.Axes[1] = NewAxis(mv.Bounds.Height)
	mv.Refresh(now, bounds)
	mv.Record = model.NewRecord("mover", int32(categories.Mover), &mv.MoverItem, model.JSON)
	return mv
}

func (mv *Mover) AddDrawer(dr Drawer) {
	mv.drawer = dr
	mv.Source = dr.Bounds()
	mv.Refresh(rl.GetTime(), mv.Bounds)
}

func (mv *Mover) GetDrawer() Drawer        { return mv.drawer }
func (mv *Mover) GetRecord() *model.Record { return mv.Record }
func (mv *Mover) GetItem() any             { return &mv.MoverItem }

func (mv *Mover) Act(can_move bool, now float64) {
	x, y := mv.Axes[0], mv.Axes[1]
	mv.drawer.Draw(rl.Vector3{X: float32(mv.Bounds.X + x.Position()),
		Y: float32(mv.Bounds.Y + y.Position()),
		Z: float32(mv.Rotation)})

	m := try.As[float64](can_move)
	y.Move(now, mv.PixelRateY*m)

	p := x.Position()
	p -= x.Move(now, mv.PixelRateX*m)
	mv.Rotation += mv.RotationRate * float32(p)
}

func (mv *Mover) Refresh(now float64, bounds rl.RectangleInt32) {
	mv.Bounds = bounds
	mv.Bounds.X += mv.Source.Width >> 1
	mv.Bounds.Y += mv.Source.Height >> 1
	mv.Bounds.Width -= mv.Source.Width
	mv.Bounds.Height -= mv.Source.Height
	mv.Axes[0].Refresh(now, mv.Bounds.Width)
	mv.Axes[1].Refresh(now, mv.Bounds.Height)
}

func (mv *Mover) Link(recs ...*model.Record) {
	err := MakeLink(mv.AddDrawer, 1, 1, recs...)
	if err != nil {
		log.Fatal(err)
	}
}

func (mv *Mover) Children() []model.Recorder {
	return []model.Recorder{mv.drawer}
}

func (mv *Mover) SetPixelRate(pixelRateX, pixelRateY float64) {
	mv.PixelRateY = pixelRateY
	mv.PixelRateX = pixelRateX
}

func (mv *Mover) GetPixelRate() (float64, float64) {
	return mv.PixelRateX, mv.PixelRateY
}
