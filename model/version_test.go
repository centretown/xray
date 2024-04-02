package model

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	list := []Version{
		{0, 0, 0, 0, time.Now(), time.Now(), IN_HOUSE},
		{0, 1, 0, 0, time.Now(), time.Now(), IN_HOUSE},
		{0, 1, 1, 0, time.Now(), time.Now(), IN_HOUSE},
		{0, 1, 1, 1, time.Now(), time.Now(), IN_HOUSE},
		{1, 1, 1, 1, time.Now(), time.Now(), IN_HOUSE},
		{1, 2, 1, 1, time.Now(), time.Now(), IN_HOUSE},
	}

	ids := make([]uint64, len(list))

	for i, v := range list {
		ids[i] = v.ToUint64()
	}

	var ver Version
	for i, id := range ids {
		lsv := list[i]
		ver.FromUint64(id)
		t.Logf("lsv %v (0x%016x) ver %v (0x%016x)", lsv, lsv.ToUint64(), ver, id)
		if lsv.Major != ver.Major {
			t.Fatal("lsv.Major!=ver.Major")
		}
		if lsv.Minor != ver.Minor {
			t.Fatal("lsv.Minor!=ver.Minor")
		}
		if lsv.Patch != ver.Patch {
			t.Fatal("lsv.Patch!=ver.Patch")
		}
		if lsv.Extension != ver.Extension {
			t.Fatal("lsv.Extension!=ver.Extension")
		}
	}
}
