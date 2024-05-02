package notes

type Composite[T any] struct {
	NoteItem
	Custom *T
}

// To complete Note interface
// Do and Values functions are required
// cmp.CanDo = false
func SetupComposite[T any](
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
