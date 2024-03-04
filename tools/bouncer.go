package tools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ CanAnimate = (*Bouncer)(nil)

const (
	max_velocity int32 = 34
	min_velocity int32 = 0
	vrange       int32 = max_velocity - min_velocity
)

type Bouncer struct {
	rect             rl.RectangleInt32
	boundsX, boundsY int32
	color_index      int32
	x_max, y_max     int32
	velocity         int32
	dx, dy           int32
	x, y             int32
}

func NewBouncer(rect rl.RectangleInt32, boundsX, boundsY int32) *Bouncer {
	anim := &Bouncer{}
	anim.Resize(rect, boundsX, boundsY)
	return anim
}

func (anim *Bouncer) Resize(rect rl.RectangleInt32, boundsX, boundsY int32) {
	if boundsX <= 0 || boundsY <= 0 {
		panic("radius <= 0")
	}
	if rect.Width <= boundsX {
		panic("ball width <= 0")
	}
	if rect.Height <= boundsY {
		panic("ball height <= 0")
	}
	anim.rect = rect
	anim.boundsX = boundsX
	anim.boundsY = boundsY
	anim.x_max = rect.Width - boundsX
	anim.y_max = rect.Height - boundsY
	anim.x = boundsX
	anim.y = boundsY
	anim.dx = 2
	anim.dy = 1
	anim.velocity = 1
}

func (anim *Bouncer) Animate(can_move int32, dr CanDraw) {
	dr.Draw(anim.rect.X+anim.x, anim.rect.Y+anim.y)

	nextX := anim.x + anim.dx
	nextY := anim.y + anim.dy

	reverse_x := (nextX >= anim.x_max) || (nextX < anim.boundsX)
	reverse_y := (nextY >= anim.y_max) || (nextY < anim.boundsY)

	anim.dx *= B2int32(!reverse_x) - B2int32(reverse_x)
	anim.dy *= B2int32(!reverse_y) - B2int32(reverse_y)

	anim.x += anim.dx * can_move
	anim.y += anim.dy * anim.velocity * can_move
	anim.velocity = anim.y * vrange / anim.rect.Height

}
