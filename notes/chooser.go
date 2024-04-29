package notes

import (
	"fmt"
	"log"

	"github.com/centretown/xray/numbers"
)

type Chooser[T fmt.Stringer] struct {
	NoteItem
	List    *[]T
	Current int
}

func NewChooser[T fmt.Stringer](
	label string,
	format string,
	list *[]T,
) *Chooser[T] {
	cho := &Chooser[T]{
		NoteItem: NoteItem{
			LabelKey:  label,
			FormatKey: format,
			CanDo:     true,
		},
		List: list,
	}

	var _ Note = cho
	return cho
}

func (cho *Chooser[T]) Item() *NoteItem {
	return &cho.NoteItem
}

func (cho *Chooser[T]) Do(command COMMAND, args ...any) {
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

func (cho *Chooser[T]) Bind(value any) {
	if err := bind(cho.List, value); err != nil {
		log.Fatal(err)
	}
}
