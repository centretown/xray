package tools

import (
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/rayl"
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Moveable = (*Mover)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type Mover struct {
	Item   *model.Record
	Source rayl.RectangleInt32
	Bounds rayl.RectangleInt32

	PixelRateY   float64 // pixels per second
	PixelRateX   float64
	Rotation     float32
	RotationRate float32
	Actions      []Action
	drawer       Drawable
}

func NewBouncer(drawer Drawable, bounds rayl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotationRate float32) *Mover {

	bc := &Mover{
		Source:       drawer.Rect(),
		Bounds:       bounds,
		PixelRateX:   pixelRateX,
		PixelRateY:   pixelRateY,
		Rotation:     0,
		RotationRate: rotationRate,
		Actions:      make([]Action, 2),
		drawer:       drawer,
	}

	bc.adjustBounds()
	bc.Actions[0] = NewAxis(rl.GetTime(), bc.Bounds.Width)
	bc.Actions[1] = NewAxis(rl.GetTime(), bc.Bounds.Height)
	bc.Item = model.NewItem("bouncer", model.Mover, bc)
	return bc
}

func (bc *Mover) Drawer() Drawable { return bc.drawer }

func (bc *Mover) adjustBounds() {
	bc.Bounds.X += bc.Source.Width / 2
	bc.Bounds.Y += bc.Source.Height / 2
	bc.Bounds.Width -= bc.Source.Width
	bc.Bounds.Height -= bc.Source.Height
}

func (bc *Mover) SetPixelRate(pixelRateX, pixelRateY float64) {
	bc.PixelRateY = pixelRateY
	bc.PixelRateX = pixelRateX
}

func (bc *Mover) GetPixelRate() (float64, float64) {
	return bc.PixelRateX, bc.PixelRateY
}

func (bc *Mover) Refresh(now float64, bounds rayl.RectangleInt32) {
	bc.Bounds = bounds
	bc.adjustBounds()
	bc.Actions[0].Refresh(now, bounds.Width-bc.Source.Width)
	bc.Actions[1].Refresh(now, bounds.Height-bc.Source.Height)
}

func (bc *Mover) Move(can_move bool, now float64) {
	x, y := bc.Actions[0], bc.Actions[1]
	bc.drawer.Draw(rayl.Vector3{X: float32(bc.Bounds.X + x.Position()),
		Y: float32(bc.Bounds.Y + y.Position()),
		Z: float32(bc.Rotation)})

	m := try.As[float64](can_move)
	y.Next(now, bc.PixelRateY*m)

	p := x.Position()
	p -= x.Next(now, bc.PixelRateX*m)
	bc.Rotation += bc.RotationRate * float32(p)
}
