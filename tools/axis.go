package tools

import (
	"github.com/centretown/xray/try"
)

type Axis struct {
	LastTime  float64 // seconds
	Max       int32
	Position  int32
	Direction int32
}

func NewAxis(max int32, now float64) *Axis {
	return &Axis{
		Max:       max,
		LastTime:  now,
		Direction: 1,
	}
}

func (ax *Axis) Refresh(max int32, now float64) {
	ax.Max = max
	ax.LastTime = now
}

func (ax *Axis) Next(current, rate float64) {
	delta := current - ax.LastTime
	deltaPos := int32(delta * rate)
	ax.LastTime += delta * try.As[float64](deltaPos != 0)

	newPos := ax.Position + deltaPos*ax.Direction

	less := newPos < 0
	more := newPos >= ax.Max
	outside := less || more
	ax.Direction *= try.As[int32](!outside) - try.As[int32](outside)
	// fmt.Print(ax.Direction)

	// if more {
	// 	ax.Position = ax.Max - deltaPos
	// 	// ax.Direction = -1
	// } else if less {
	// 	ax.Position = -deltaPos
	// 	// ax.Direction = 1
	// } else {
	// 	ax.Position = newPos
	// }

	ax.Position = try.As[int32](more)*(ax.Max-deltaPos) +
		try.As[int32](less)*deltaPos +
		try.As[int32](!less && !more)*newPos
}
