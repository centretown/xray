package gizmo

import (
	"log"

	"github.com/centretown/xray/check"
	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Mover = (*Tracker)(nil)
var _ model.Parent = (*Tracker)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type TrackerItem struct {
	Rectangle    rl.RectangleInt32
	Source       rl.RectangleInt32
	PixelRateY   float64 // pixels per second
	PixelRateX   float64
	Rotation     float32
	RotationRate float32
	Axes         []*Axis
	drawer       Drawer
}

type Tracker struct {
	TrackerItem
	Record *model.Record
}

func NewTracker(bounds rl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotationRate float32) *Tracker {

	mv := &Tracker{}
	mv.Rectangle = bounds
	mv.PixelRateX = pixelRateX
	mv.PixelRateY = pixelRateY
	mv.Rotation = 0
	mv.RotationRate = rotationRate
	mv.Axes = make([]*Axis, 2)

	now := rl.GetTime()
	mv.Axes[0] = NewAxis(mv.Rectangle.Width)
	mv.Axes[1] = NewAxis(mv.Rectangle.Height)
	mv.Refresh(now, bounds)
	mv.Record = model.NewRecord("mover", int32(categories.Mover), &mv.TrackerItem, model.JSON)
	return mv
}

func (mv *Tracker) AddDrawer(dr Drawer) {
	mv.drawer = dr
	mv.Source = dr.Bounds()
	mv.Refresh(rl.GetTime(), mv.Rectangle)
}

func (mv *Tracker) GetDrawer() Drawer        { return mv.drawer }
func (mv *Tracker) GetRecord() *model.Record { return mv.Record }
func (mv *Tracker) GetItem() any             { return &mv.TrackerItem }

// func (mv *Tracker) Bounds() rl.RectangleInt32 { return mv.Bounds }

func (mv *Tracker) Move(can_move bool, now float64) {
	x, y := mv.Axes[0], mv.Axes[1]
	mv.drawer.Draw(rl.Vector3{X: float32(mv.Rectangle.X + x.Position()),
		Y: float32(mv.Rectangle.Y + y.Position()),
		Z: float32(mv.Rotation)})

	m := check.As[float64](can_move)
	y.Move(now, mv.PixelRateY*m)

	p := x.Position()
	p -= x.Move(now, mv.PixelRateX*m)
	mv.Rotation += mv.RotationRate * float32(p)
}

func (mv *Tracker) Refresh(now float64, bounds rl.RectangleInt32, fs ...func(any)) {
	mv.Rectangle = bounds
	mv.Rectangle.X += mv.Source.Width >> 1
	mv.Rectangle.Y += mv.Source.Height >> 1
	mv.Rectangle.Width -= mv.Source.Width
	mv.Rectangle.Height -= mv.Source.Height
	mv.Axes[0].Refresh(now, mv.Rectangle.Width)
	mv.Axes[1].Refresh(now, mv.Rectangle.Height)
}

func (mv *Tracker) LinkChildren(recs ...*model.Record) {
	err := MakeLink(mv.AddDrawer, 1, 1, recs...)
	if err != nil {
		log.Fatal(err)
	}
}

func (mv *Tracker) Children() []model.Recorder {
	return []model.Recorder{mv.drawer}
}

func (mv *Tracker) SetPixelRate(pixelRateX, pixelRateY float64) {
	mv.PixelRateY = pixelRateY
	mv.PixelRateX = pixelRateX
}

func (mv *Tracker) GetPixelRate() (float64, float64) {
	return mv.PixelRateX, mv.PixelRateY
}
