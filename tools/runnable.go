package tools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Runnable struct {
	draw  Drawable
	anim  Animator
	next  float64
	after float64
}

func NewRunnable(draw Drawable, anim Animator, start float64, after float64) *Runnable {
	run := &Runnable{
		draw:  draw,
		anim:  anim,
		next:  start + after,
		after: after,
	}
	return run
}

func (run *Runnable) Animate(can_move int32, current float64) {
	can_move = can_move * B2int32(current >= run.next)
	run.anim.Animate(can_move, run.draw)
}

func (run *Runnable) Resize(rect rl.RectangleInt32, boundsX int32, boundsY int32, current float64) {
	run.anim.Resize(rect, boundsX, boundsY)
	run.next = current + run.after
}
