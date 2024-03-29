package tools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor struct {
	draw  Drawable
	anim  Moveable
	next  float64
	after float64
}

func NewActor(draw Drawable, anim Moveable, after float64) *Actor {
	act := &Actor{
		draw:  draw,
		anim:  anim,
		next:  0,
		after: after,
	}
	return act
}

func (act *Actor) Animate(can_move bool, current float64) {
	act.anim.Draw(can_move, current, act.draw)
}

func (act *Actor) Resize(rect rl.RectangleInt32, current float64) {
	act.anim.Refresh(current, rect)
	act.next = current + act.after
}
