package notes

type Doer interface {
	Values() []any
	Do(command COMMAND, args ...any)
}

type Output struct {
	Label string
	Value string
}

type NoteItem struct {
	LabelKey  string
	FormatKey string
	Output    Output
	CanDo     bool
}

type Note interface {
	Doer
	Item() *NoteItem
}
