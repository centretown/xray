package dbg

import (
	"github.com/centretown/xray/model"
)

type Schema struct {
	Version model.Version
	Create  []string

	InsertVersion string
	InsertItem    string
	InsertLink    string
	InsertTag     string

	GetItem  string
	GetLink  string
	GetLinks string
}

func NewSchema() *Schema {
	return SchemaGame
}
