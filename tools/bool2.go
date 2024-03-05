package tools

// The compiler currently only optimizes this form.
//
// See issue 6011.
//
// go build -gcflags='-l -v' tb.go
// go tool objdump -S -s B2I tb
// code is optimized to:  return i
//
//	asm: MOVZX AL, AX
//	     RET
//
// use for branch free conditional
// hopefully this gets inlined

func Bool2int(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func Bool2int32(b bool) int32 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return int32(i)
}

func Bool2float32(b bool) float32 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return float32(i)
}
