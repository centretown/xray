package dbio

import (
	"github.com/centretown/xray/model"
)

type Schema struct {
	Version       model.Version
	Create        []string
	InsertVersion string
	InsertItem    string
}

func NewSchema() *Schema {
	return SchemaGame
}
