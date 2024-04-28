package notes

import "log"

type Composite[T any] struct {
	NoteItem
	Custom *T
}

// To complete Note interface
// Do and Values functions are required
// cmp.CanDo = false
func InitComposite[T any](
	cmp *Composite[T],
	label string,
	format string,
	custom *T,
) {
	cmp.LabelKey = label
	cmp.FormatKey = format
	cmp.Custom = custom
}

func (cmp *Composite[T]) Item() *NoteItem {
	return &cmp.NoteItem
}

func (cmp *Composite[T]) Bind(value any) {
	if err := bind(cmp.Custom, value); err != nil {
		log.Fatal(err)
	}
}
