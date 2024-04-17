package gizmo

import (
	"github.com/centretown/xray/check"
)

// var _ Motor = (*Axis)(nil)

type Axis struct {
	Position  int32
	Direction int32
	Last      float64 // seconds
	Max       int32
}

func NewAxis(max int32) *Axis {
	return &Axis{
		Max:       max,
		Direction: 1,
	}
}

func (ax *Axis) Refresh(now float64, max int32) {
	ax.Last = now
	ax.Max = max
}

func (ax *Axis) Move(current, rate float64) int32 {
	var (
		delta    = current - ax.Last
		deltaPos = int32(delta * rate)
		newPos   = ax.Position + deltaPos*ax.Direction
		less     = newPos < 0
		more     = newPos >= ax.Max
		outside  = less || more
	)

	ax.Last += delta * check.As[float64](deltaPos != 0)
	ax.Direction *= check.As[int32](!outside) - check.As[int32](outside)
	ax.Position = check.As[int32](more)*(ax.Max-deltaPos) +
		check.As[int32](less)*deltaPos +
		check.As[int32](!outside)*newPos
	return ax.Position
}
