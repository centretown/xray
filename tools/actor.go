package tools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Actor struct {
	Character Drawable
	Mover     Moveable
}

func NewActor(draw Drawable, anim Moveable, after float64) *Actor {
	act := &Actor{
		Character: draw,
		Mover:     anim,
	}
	return act
}

func (act *Actor) Animate(can_move bool, current float64) {
	act.Mover.Draw(can_move, current, act.Character)
}

func (act *Actor) Resize(rect rl.RectangleInt32, current float64) {
	act.Mover.Refresh(current, rect)
}
