package model

type Category int

const (
	Picture Category = iota
	Bouncer          //Pink Elephant
	Scene
	Person
	Ball
)

//go:generate stringer -type=Category
