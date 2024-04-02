package access

import (
	"testing"

	"github.com/google/uuid"
)

func TestUuid(t *testing.T) {
	m := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		id, err := uuid.NewV7()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id.String(), len(id.String()))
		m[id.String()] = false
	}

	if len(m) != 1000 {
		t.Fatalf("1000 expected %d received", len(m))
	}
}
