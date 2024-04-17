package categories

type Category int32

type NotFound struct{}

const (
	Unknown Category = iota
	Game
	Ellipse
	Texture
	Tracker
	LifeMover
	LifeGrid
	Player
	CellsOrg
	COUNT
)

//go:generate stringer -type=Category
