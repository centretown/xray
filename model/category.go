package model

type Category int

const (
	Picture Category = iota
	Mover            //Pink Elephant
	Scene
	Person
	Ball
)

//go:generate stringer -type=Category
