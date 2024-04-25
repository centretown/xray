package gizzmo

import "github.com/centretown/xray/numbers"

type NumberCell[T numbers.NumberType] struct {
	States []T
}

func NewNumberCell[T numbers.NumberType](states int) *NumberCell[T] {

	nc := &NumberCell[T]{
		States: make([]T, states),
	}
	return nc
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
