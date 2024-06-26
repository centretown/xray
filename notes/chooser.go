package notes

import (
	"fmt"

	"github.com/centretown/xray/numbers"
)

type Chooser[T fmt.Stringer] struct {
	Scribe
	List    *[]T
	Current int
}

func NewChooser[T fmt.Stringer](
	label string,
	format string,
	list *[]T,
) *Chooser[T] {
	cho := &Chooser[T]{
		Scribe: Scribe{
			LabelKey:  label,
			FormatKey: format,
			CanDo:     true,
		},
		List: list,
	}

	var _ Note = cho
	return cho
}

func (cho *Chooser[T]) GetScribe() *Scribe {
	return &cho.Scribe
}

func (cho *Chooser[T]) Do(command Command, args ...any) {
	var (
		length    = len(*cho.List)
		selection int
	)

	switch command {
	case SET:
		if len(args) > 0 {
			selection = args[0].(int)
		}
		cho.Current = numbers.AsOr(selection >= 0 && selection < length, selection, 0)
	case INCREMENT_MORE, INCREMENT:
		cho.Current = numbers.AsOr(cho.Current+1 < length, cho.Current+1, 0)
	case DECREMENT_MORE, DECREMENT:
		cho.Current = numbers.AsOr(cho.Current-1 >= 0, cho.Current-1, length-1)
	}
}

func (cho *Chooser[T]) Values() []any {
	return []any{(*cho.List)[cho.Current].String()}
}
