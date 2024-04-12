package gizmo

type CellState uint64

type Cell[T any] struct {
	States []T
}

func (c *Cell[T]) GetState(i int) T {
	var out T
	if i < len(c.States) {
		return c.States[i]
	}
	return out
}

func (c *Cell[T]) SetState(i int, v T) {
	if i < len(c.States) {
		c.States[i] = v
	}
}

func (c *Cell[T]) Clear() {
	var t T
	for i := range c.States {
		c.States[i] = t
	}
}
