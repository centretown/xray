package notes

import (
	"github.com/centretown/xray/numbers"
)

type Ranger[T numbers.NumberType] struct {
	Scribe
	Min     T
	Max     T
	More    T
	Current *T
}

func NewRanger[T numbers.NumberType](
	label string,
	format string,
	current *T,
	min T,
	max T,
	more T,
) *Ranger[T] {
	rngr := &Ranger[T]{
		Scribe: Scribe{
			LabelKey:  label,
			FormatKey: format,
			CanDo:     true,
		},
		Current: current,
		Min:     min,
		Max:     max,
		More:    more,
	}
	// var _ Note = rngr
	return rngr
}

func (rngr *Ranger[T]) GetScribe() *Scribe {
	return &rngr.Scribe
}

func (rngr *Ranger[T]) Do(command Command, args ...any) {
	var (
		current   = *rngr.Current
		next      T
		selection = rngr.Min
	)
	current = numbers.AsOr(current < rngr.Min, rngr.Min, current)
	current = numbers.AsOr(current >= rngr.Max, rngr.Max-1, current)

	switch command {
	case SET:
		if len(args) > 0 {
			selection = args[0].(T)
		}
		current = numbers.AsOr(selection >= rngr.Min && selection < rngr.Max,
			selection, rngr.Min)
	case INCREMENT_MORE:
		next = current + rngr.More
		current = numbers.AsOr(next < rngr.Max, next, rngr.Max-1)
	case INCREMENT:
		next = current + 1
		current = numbers.AsOr(next < rngr.Max, next, rngr.Max-1)
	case DECREMENT_MORE:
		next = current - rngr.More
		current = numbers.AsOr(next >= rngr.Min, next, rngr.Min)
	case DECREMENT:
		next = current - 1
		current = numbers.AsOr(next >= rngr.Min, next, rngr.Min)
	}
	*rngr.Current = current
}

// func (rngr *Ranger[T]) Values() []any {
// 	return []any{*rngr.Current}
// }
