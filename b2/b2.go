package b2

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
// hopefully this gets inlined (so far it does)

func To[T int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](b bool) T {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return T(i)
}

func ToInt(b bool) int {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

func ToInt32(b bool) int32 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return int32(i)
}

func ToInt64(b bool) uint64 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return uint64(i)
}

func ToFloat32(b bool) float32 {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return float32(i)
}
