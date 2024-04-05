package model

import (
	"testing"
)

func TestSchema(t *testing.T) {
	f := NewRecord("Dave", 4, "had a little lamb")
	t.Log(f.Title, f.Category, f.Major, f.Minor, f.OriginMajor, f.OriginMinor, f.Content)

}
