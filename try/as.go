package try

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

func As[T int | uint | int8 | int16 | int32 | int64 |
	uint8 | uint16 | uint32 | uint64 |
	float32 | float64](condition bool) T {

	var i int
	if condition {
		i = 1
	} else {
		i = 0
	}
	return T(i)
}

func Or[T int | uint | int8 | int16 | int32 | int64 |
	uint8 | uint16 | uint32 | uint64 |
	float32 | float64](condition bool, falseVal, trueVal T) T {
	return falseVal + (trueVal-falseVal)*As[T](condition)
}
