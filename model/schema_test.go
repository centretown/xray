package model

import (
	"testing"
)

func TestSchema(t *testing.T) {
	f := NewItem("Dave", Person, "had a little lamb")
	t.Log(f.Title, f.Category, f.ID, f.Origin, f.Content)

}
