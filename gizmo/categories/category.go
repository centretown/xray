package categories

type Category int32

type NotFound struct{}

const (
	Texture Category = iota
	Circle
	CellsOrg
	Mover
	Game
	NumberMoveri8
	NumberGridi8
	Player
	COUNT
)

//go:generate stringer -type=Category
