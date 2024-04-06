package gizmo

type Category int32

const (
// Texture Category = iota
// Circle
// Motor
// Game
// Player
)

//go:generate stringer -type=Category
