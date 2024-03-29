package tools

import (
	"github.com/centretown/xray/b2"
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
	ax.LastTime += delta * b2.To[float64](deltaPos != 0)

	newPos := ax.Position + deltaPos*ax.Direction

	less := newPos < 0
	more := newPos >= ax.Max
	outside := less || more
	ax.Direction *= b2.To[int32](!outside) - b2.To[int32](outside)
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

	ax.Position = b2.To[int32](more)*(ax.Max-deltaPos) +
		b2.To[int32](less)*-deltaPos +
		b2.To[int32](!less && !more)*newPos
}
