package tools

import (
	"xray/b2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor struct {
	draw  CanDraw
	anim  CanAnimate
	next  float64
	after float64
}

func NewActor(draw CanDraw, anim CanAnimate, after float64) *Actor {
	act := &Actor{
		draw:  draw,
		anim:  anim,
		next:  0,
		after: after,
	}
	return act
}

func (act *Actor) Animate(can_move int32, current float64) {
	can_move = can_move * b2.ToInt32(current >= act.next)
	act.anim.Animate(can_move, act.draw)
}

func (act *Actor) Resize(rect rl.RectangleInt32, current float64) {
	act.anim.Resize(rect, act.draw.Width(), act.draw.Height())
	act.next = current + act.after
}
