package gizmo

import (
	"github.com/centretown/xray/try"
)

var _ Motor = (*Axis)(nil)

type Axis struct {
	Pos  int32
	Dir  int32
	Last float64 // seconds
	Max  int32
}

func NewAxis(now float64, max int32) *Axis {
	return &Axis{
		Last: now,
		Max:  max,
		Dir:  1,
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
		newPos   = ax.Pos + deltaPos*ax.Dir
		less     = newPos < 0
		more     = newPos >= ax.Max
		outside  = less || more
	)

	ax.Last += delta * try.As[float64](deltaPos != 0)
	ax.Dir *= try.As[int32](!outside) - try.As[int32](outside)
	ax.Pos = try.As[int32](more)*(ax.Max-deltaPos) +
		try.As[int32](less)*deltaPos +
		try.As[int32](!outside)*newPos
	return ax.Pos
}

func (ax *Axis) Position() int32  { return ax.Pos }
func (ax *Axis) Direction() int32 { return ax.Dir }
