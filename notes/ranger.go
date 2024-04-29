package notes

import (
	"log"

	"github.com/centretown/xray/numbers"
)

type Ranger[T numbers.NumberType] struct {
	NoteItem
	min     T
	max     T
	more    T
	current *T
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
		NoteItem: NoteItem{
			LabelKey:  label,
			FormatKey: format,
			CanDo:     true,
		},
		current: current,
		min:     min,
		max:     max,
		more:    more,
	}
	var _ Note = rngr
	return rngr
}

func (rngr *Ranger[T]) Item() *NoteItem {
	return &rngr.NoteItem
}

func (rngr *Ranger[T]) Do(command COMMAND, args ...any) {
	var (
		current   = *rngr.current
		next      T
		selection = rngr.min
	)
	current = numbers.AsOr(current < rngr.min, rngr.min, current)
	current = numbers.AsOr(current >= rngr.max, rngr.max-1, current)

	switch command {
	case SET:
		if len(args) > 0 {
			selection = args[0].(T)
		}
		current = numbers.AsOr(selection >= rngr.min && selection < rngr.max,
			selection, rngr.min)
	case INCREMENT_MORE:
		next = current + rngr.more
		current = numbers.AsOr(next < rngr.max, next, rngr.max-1)
	case INCREMENT:
		next = current + 1
		current = numbers.AsOr(next < rngr.max, next, rngr.max-1)
	case DECREMENT_MORE:
		next = current - rngr.more
		current = numbers.AsOr(next >= rngr.min, next, rngr.min)
	case DECREMENT:
		next = current - 1
		current = numbers.AsOr(next >= rngr.min, next, rngr.min)
	}
	*rngr.current = current
}

func (rngr *Ranger[T]) Values() []any {
	return []any{*rngr.current}
}

func (rngr *Ranger[T]) Bind(value any) {
	if err := bind(rngr.current, value); err != nil {
		log.Fatal(err)
	}
}
