package gizmo

import (
	"encoding/json"
	"fmt"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Actor = (*Mover)(nil)
var _ model.Recorder = (*Mover)(nil)
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

	mv.adjustBounds()
	mv.Axes[0] = NewAxis(rl.GetTime(), mv.Bounds.Width)
	mv.Axes[1] = NewAxis(rl.GetTime(), mv.Bounds.Height)
	mv.Record = model.NewRecord("mover", int32(categories.Mover), &mv.MoverItem, model.JSON)
	return mv
}
func (mv *Mover) AddDrawer(dr Drawer) *Mover {
	mv.drawer = dr
	mv.Source = dr.Bounds()
	return mv
}

func (mv *Mover) GetDrawer() Drawer        { return mv.drawer }
func (mv *Mover) GetRecord() *model.Record { return mv.Record }
func (mv *Mover) GetItem() any             { return &mv.MoverItem }

func (mv *Mover) Decode(rec *model.Record) (err error) {
	mv.Record = rec
	cat := categories.Category(rec.Category)
	if cat == categories.Mover {
		err = json.Unmarshal([]byte(rec.Content), &mv.MoverItem)
	} else {
		err = fmt.Errorf("wrong category want %s have %s",
			categories.Mover, cat)
	}

	return
}

func (mv *Mover) adjustBounds() {
	mv.Bounds.X += mv.Source.Width / 2
	mv.Bounds.Y += mv.Source.Height / 2
	mv.Bounds.Width -= mv.Source.Width
	mv.Bounds.Height -= mv.Source.Height
}

func (mv *Mover) SetPixelRate(pixelRateX, pixelRateY float64) {
	mv.PixelRateY = pixelRateY
	mv.PixelRateX = pixelRateX
}

func (mv *Mover) GetPixelRate() (float64, float64) {
	return mv.PixelRateX, mv.PixelRateY
}

func (mv *Mover) Refresh(now float64, bounds rl.RectangleInt32) {
	mv.Bounds = bounds
	mv.adjustBounds()
	mv.Axes[0].Refresh(now, bounds.Width-mv.Source.Width)
	mv.Axes[1].Refresh(now, bounds.Height-mv.Source.Height)
}

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

func (mv *Mover) Children() (rcds []model.Recorder) {
	rcds = make([]model.Recorder, 0)
	r, ok := mv.drawer.(model.Recorder)
	if ok {
		rcds = append(rcds, r)
	}
	return
}

func (mv *Mover) Link(recs ...*model.Record) {
	if len(recs) < 1 {
		return
	}
	if len(recs) > 1 {
		fmt.Println("Mover.Link too many links want 1 have", len(recs))
	}
	var err error

	rec := recs[0]
	cat := categories.Category(rec.Category)
	if cat == categories.Circle {
		circle := &Circle{}
		err = circle.Decode(rec)
		mv.AddDrawer(circle)
	} else if cat == categories.Texture {
		tex := &Texture{}
		err = tex.Decode(rec)
		mv.AddDrawer(tex)
	} else {
		err = fmt.Errorf("wrong category want %s or %s have %s",
			categories.Circle, categories.Texture, cat)
	}

	if err != nil {
		fmt.Println(err)
	}
}
