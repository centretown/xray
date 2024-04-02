package model

type Scheme int

const (
	None Scheme = iota // none
	File               // file
)

//go:generate stringer -linecomment -type=Scheme
