package categories

type Category int32

type NotFound struct{}

const (
	Texture Category = iota
	Circle
	CellsOrg
	Mover
	Game
	CellsMover
	Grid
	Player
	COUNT
)

//go:generate stringer -type=Category
