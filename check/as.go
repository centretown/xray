// Package try provides useful tools for branchless programming
//
// The compiler currently assembles this form:
//
//	func ToInt(b bool) int {
//		var i int
//		if b {
//			i = 1
//		} else {
//			i = 0
//		}
//		return i
//	}
//
// To ths:
//
//	asm: MOVZX AL, AX
//
// See issue 6011.
// The As function is inlined.
package check

import "golang.org/x/exp/constraints"

// NumberType is a constraint for all values that can be set to one or zero
type NumberType interface {
	constraints.Integer | constraints.Float
}

// Branchless way to get 1 or 0
func As[T NumberType](condition bool) T {
	var i int
	if condition {
		i = 1
	} else {
		i = 0
	}
	return T(i)
}

// Branchless way to get one of 2 values (that are not 1 or 0)
func AsOr[T NumberType](condition bool, trueVal, falseVal T) T {
	return falseVal + (trueVal-falseVal)*As[T](condition)
}

func Is[T NumberType](value T) bool {
	var none T
	return value != none
}

func Or[T any](condition bool, vals [2]T) T {
	return vals[As[int](!condition)]
}
