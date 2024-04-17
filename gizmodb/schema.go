package gizmodb

import (
	"github.com/centretown/xray/gizmodb/model"
)

type Schema struct {
	Version model.Version
	Create  []string

	InsertVersion string
	InsertItem    string
	InsertLink    string
	InsertTag     string

	GetVersions string
	GetVersion  string
	GetItem     string
	GetLink     string
	GetLinks    string
}

func NewSchema() *Schema {
	return SchemaGame
}
