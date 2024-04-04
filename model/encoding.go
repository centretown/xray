package model

type Encoding int32

const (
	RAW Encoding = iota
	JSON
	YAML
	XML
	CSV
)

//go:generate stringer -type=Encoding
