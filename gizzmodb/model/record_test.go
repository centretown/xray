package model

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
		m[id.String()] = false
		major, minor := RecordID(id)
		idd := RecordUUID(major, minor)
		if idd != id {
			t.Logf("%x", id)
			t.Logf("%x", idd)
			t.Fatalf("major%x minor%x", major, minor)
		}

		t.Logf("%x %x %s,%d", major, minor, id.String(), len(id.String()))
	}

	if len(m) != 1000 {
		t.Fatalf("1000 expected %d received", len(m))
	}
}
