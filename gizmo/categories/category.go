package categories

type Category int32

type NotFound struct{}

const (
	Texture Category = iota
	Circle
	Mover
	Game
	Player
	COUNT
)

//go:generate stringer -type=Category
