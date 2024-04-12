package gizmo

import "github.com/centretown/xray/check"

type NumberCell[T check.NumberType] struct {
	States []T
}

func NewNumberCell[T check.NumberType](states int) *NumberCell[T] {

	nc := &NumberCell[T]{
		States: make([]T, states),
	}
	return nc
}

func Zero[T any]() (t T) {
	return t
}

func (c *NumberCell[T]) Clear() {
	var t T
	for i := range c.States {
		c.States[i] = t
	}
}

func (c *NumberCell[T]) Get(i int32) T {
	return c.States[i]
}

func (c *NumberCell[T]) Set(i int32, v T) {
	c.States[i] = v
}
