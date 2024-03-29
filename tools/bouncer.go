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
	xAxis        *Axis
	yAxis        *Axis
}

func NewBouncer(source, bounds rl.RectangleInt32,
	pixelRateX, pixelRateY float64, rotation float32) *Bouncer {

	bc := &Bouncer{
		source:       source,
		bounds:       bounds,
		pixelRateX:   pixelRateX,
		pixelRateY:   pixelRateY,
		rotationRate: rotation,
	}

	bc.adjustBounds()

	bc.xAxis = NewAxis(bc.bounds.Width, rl.GetTime())
	bc.yAxis = NewAxis(bc.bounds.Height, rl.GetTime())
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

func (bc *Bouncer) Refresh(current float64, bounds rl.RectangleInt32) {
	bc.bounds = bounds
	bc.adjustBounds()
	bc.xAxis.Refresh(bounds.Width-bc.source.Width, current)
	bc.yAxis.Refresh(bounds.Height-bc.source.Height, current)
}

func (bc *Bouncer) Draw(can_move bool, current float64, dr Drawable) {
	x, y := bc.xAxis, bc.yAxis
	dr.Draw(rl.Vector3{X: float32(bc.bounds.X + x.Position),
		Y: float32(bc.bounds.Y + y.Position),
		Z: float32(bc.rotation)})

	m := try.As[float64](can_move)
	x.Next(current, bc.pixelRateX*m)
	// gRateY := float64(bc.bounds.Y) + bc.pixelRateY*
	// 	(float64(y.Position)/
	// 		(float64(y.Max)/2))
	y.Next(current, bc.pixelRateY*m)
	bc.rotation += bc.rotationRate
}
