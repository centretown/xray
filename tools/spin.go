package tools

var _ Action = (*Axis)(nil)

type Spin struct {
	LastTime  float64 // seconds
	Direction int32
}
