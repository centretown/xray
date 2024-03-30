package tools

var _ Action = (*Spin)(nil)

type Spin struct {
	rotation float32
	rate     float32
	master   Action
	position int32
}

func NewSpin(action Action, rate float32) *Spin {
	return &Spin{
		master:   action,
		rate:     rate,
		position: action.Position(),
	}
}

func (sp *Spin) Position() int32 {
	return sp.master.Position()
}

func (sp *Spin) Rotation() float32 {
	return sp.rotation
}

func (sp *Spin) Direction() int32 {
	return sp.master.Direction()
}

func (sp *Spin) Refresh(now float64, position, max int32) {
}

func (sp *Spin) Next(current, rate float64) (position int32) {
	sp.position -= sp.master.Position()
	sp.rotation += sp.rate * float32(sp.position)
	return sp.position
}
