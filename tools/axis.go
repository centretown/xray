package tools

import (
	"github.com/centretown/xray/try"
)

var _ Action = (*Axis)(nil)

type Axis struct {
	LastTime  float64 // seconds
	Max       int32
	position  int32
	direction int32
}

func NewAxis(max int32, now float64) *Axis {
	return &Axis{
		Max:       max,
		LastTime:  now,
		direction: 1,
	}
}

func (ax *Axis) Position() int32 {
	return ax.position
}

func (ax *Axis) Direction() int32 {
	return ax.direction
}

func (ax *Axis) Refresh(now float64, position, max int32) {
	ax.Max = max
	ax.LastTime = now
	ax.position = position
}

func (ax *Axis) Next(current, rate float64) int32 {
	var (
		delta    = current - ax.LastTime
		deltaPos = int32(delta * rate)
		newPos   = ax.position + deltaPos*ax.direction
		less     = newPos < 0
		more     = newPos >= ax.Max
		outside  = less || more
	)

	ax.LastTime += delta * try.As[float64](deltaPos != 0)
	ax.direction *= try.As[int32](!outside) - try.As[int32](outside)
	ax.position = try.As[int32](more)*(ax.Max-deltaPos) +
		try.As[int32](less)*deltaPos +
		try.As[int32](!outside)*newPos
	return ax.position
}
