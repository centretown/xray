package model

type Category int

const (
	Picture Category = iota
	Mover            //Pink Elephant
	Scene
	Person
	Circle
)

//go:generate stringer -type=Category
