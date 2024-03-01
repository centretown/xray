package tools

// See issue 6011.
func B2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func B2int32(b bool) int32 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return int32(i)
}
