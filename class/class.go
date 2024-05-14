package class

type Class int32

type NotFound struct{}

const (
	ClassBegin Class = iota
	Game
	Ellipse
	Texture
	Tracker
	LifeMover
	LifeGrid
	Player
	CellsOrg
	Pane
	Vocabulary
	ClassEnd
)

//go:generate stringer -type=Class

type Genre int32

const (
	GenreBegin Genre = iota
	Shape
	NumberGrid
	GridMover
	GenreEnd
)

//go:generate stringer -type=Genre
