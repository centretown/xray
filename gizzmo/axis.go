package gizzmo

import (
	"github.com/centretown/xray/check"
)

// var _ Motor = (*Axis)(nil)

type Axis struct {
	Position  float32
	Direction float32
	Last      float64 // seconds
	Extent    float32
}

func (ax *Axis) Refresh(now float64, max float32) {
	ax.Last = now
	ax.Extent = max
}

func (ax *Axis) Move(current, rate float64) float32 {
	var (
		delta    = current - ax.Last
		deltaPos = delta * rate
		newPos   = ax.Position + float32(deltaPos)*ax.Direction
		less     = newPos < 0
		more     = newPos >= ax.Extent
		outside  = less || more
	)

	ax.Last += delta * check.As[float64](deltaPos != 0)
	ax.Direction *= check.As[float32](!outside) - check.As[float32](outside)
	ax.Position = check.As[float32](more)*(ax.Extent-float32(deltaPos)) +
		check.As[float32](less)*float32(deltaPos) +
		check.As[float32](!outside)*newPos
	return ax.Position
}
