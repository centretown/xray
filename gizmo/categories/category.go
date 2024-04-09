package categories

type Category int32

type NotFound struct{}

const (
	Texture Category = iota
	Circle
	Cells
	Mover
	Game
	CellsMover
	Player
	COUNT
)

//go:generate stringer -type=Category
