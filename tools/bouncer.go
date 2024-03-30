package tools

import (
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Moveable = (*Bouncer)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type Bouncer struct {
	source rl.RectangleInt32
	bounds rl.RectangleInt32

	pixelRateY   float64 // pixels per second
	pixelRateX   float64
	rotation     float32
	rotationRate float32
	actions      []Action
}

func NewBouncer(source, bounds rl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotationRate float32) *Bouncer {

	bc := &Bouncer{
		source:       source,
		bounds:       bounds,
		pixelRateX:   pixelRateX,
		pixelRateY:   pixelRateY,
		rotation:     0,
		rotationRate: rotationRate,
		actions:      make([]Action, 2),
	}

	bc.adjustBounds()

	bc.actions[0] = NewAxis(bc.bounds.Width, rl.GetTime())
	bc.actions[1] = NewAxis(bc.bounds.Height, rl.GetTime())
	return bc
}

func (bc *Bouncer) adjustBounds() {
	bc.bounds.X += bc.source.Width / 2
	bc.bounds.Y += bc.source.Height / 2
	bc.bounds.Width -= bc.source.Width
	bc.bounds.Height -= bc.source.Height
}

func (bc *Bouncer) SetPixelRate(pixelRateX, pixelRateY float64) {
	bc.pixelRateY = pixelRateY
	bc.pixelRateX = pixelRateX
}

func (bc *Bouncer) GetPixelRate() (float64, float64) {
	return bc.pixelRateX, bc.pixelRateY
}

func (bc *Bouncer) Refresh(now float64, bounds rl.RectangleInt32) {
	bc.bounds = bounds
	bc.adjustBounds()
	bc.actions[0].Refresh(now, 0, bounds.Width-bc.source.Width)
	// bc.actions[0].Position())
	bc.actions[1].Refresh(now, 0, bounds.Height-bc.source.Height)
	// bc.actions[1].Position())
}

func (bc *Bouncer) Draw(can_move bool, now float64, dr Drawable) {
	x, y := bc.actions[0], bc.actions[1]
	dr.Draw(rl.Vector3{X: float32(bc.bounds.X + x.Position()),
		Y: float32(bc.bounds.Y + y.Position()),
		Z: float32(bc.rotation)})

	m := try.As[float64](can_move)
	y.Next(now, bc.pixelRateY*m)

	p := x.Position()
	p -= x.Next(now, bc.pixelRateX*m)
	bc.rotation += bc.rotationRate * float32(p)
}
