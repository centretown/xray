package dbio

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
}

func NewSchema() *Schema {
	return SchemaGame
}
