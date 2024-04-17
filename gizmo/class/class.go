package class

type Class int32

type NotFound struct{}

const (
	Unknown Class = iota
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

//go:generate stringer -type=Class
