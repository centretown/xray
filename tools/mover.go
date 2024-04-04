package tools

import (
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Moveable = (*Mover)(nil)

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
	Actions      []Action
	drawer       Drawable
}

type Mover struct {
	MoverItem
	Record *model.Record
}

func NewMover(drawer Drawable, bounds rl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotationRate float32) *Mover {

	mv := &Mover{}
	mv.Source = drawer.Rect()
	mv.Bounds = bounds
	mv.PixelRateX = pixelRateX
	mv.PixelRateY = pixelRateY
	mv.Rotation = 0
	mv.RotationRate = rotationRate
	mv.Actions = make([]Action, 2)
	mv.drawer = drawer

	mv.adjustBounds()
	mv.Actions[0] = NewAxis(rl.GetTime(), mv.Bounds.Width)
	mv.Actions[1] = NewAxis(rl.GetTime(), mv.Bounds.Height)
	mv.Record = model.NewRecord("mover", model.Mover, &mv.MoverItem)
	return mv
}

func (mv *Mover) Drawer() Drawable         { return mv.drawer }
func (mv *Mover) GetRecord() *model.Record { return mv.Record }

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
	mv.Actions[0].Refresh(now, bounds.Width-mv.Source.Width)
	mv.Actions[1].Refresh(now, bounds.Height-mv.Source.Height)
}

func (mv *Mover) Move(can_move bool, now float64) {
	x, y := mv.Actions[0], mv.Actions[1]
	mv.drawer.Draw(rl.Vector3{X: float32(mv.Bounds.X + x.Position()),
		Y: float32(mv.Bounds.Y + y.Position()),
		Z: float32(mv.Rotation)})

	m := try.As[float64](can_move)
	y.Next(now, mv.PixelRateY*m)

	p := x.Position()
	p -= x.Next(now, mv.PixelRateX*m)
	mv.Rotation += mv.RotationRate * float32(p)
}
