package gizzmo

import (
	"github.com/centretown/xray/check"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// var _ Motor = (*Axis)(nil)

type Axis struct {
	Position   float32
	Direction  float32
	LastUpdate float64 // seconds
	Extent     rl.Vector2
}

func (ax *Axis) Setup(now float64, min, max float32) {
	ax.LastUpdate = now
	ax.Extent = rl.Vector2{X: min, Y: max}
	ax.Position = min
	ax.Direction = 1
}

func (ax *Axis) Refresh(now float64, extent rl.Vector2) {
	ax.LastUpdate = now
	ax.Extent = extent
}

func (ax *Axis) Move(current, rate float64) float32 {
	var (
		delta    = current - ax.LastUpdate
		deltaPos = delta * rate
		newPos   = ax.Position + float32(deltaPos)*ax.Direction
		less     = newPos < ax.Extent.X
		more     = newPos >= ax.Extent.Y
		outside  = less || more
	)

	ax.LastUpdate += delta * check.As[float64](deltaPos != 0)
	ax.Direction *= check.As[float32](!outside) - check.As[float32](outside)
	ax.Position = check.As[float32](more)*(ax.Extent.Y-float32(deltaPos)) +
		check.As[float32](less)*(ax.Extent.X+float32(deltaPos)) +
		check.As[float32](!outside)*newPos
	return ax.Position
}
