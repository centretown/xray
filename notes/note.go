package notes

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
	Values() []any
	Do(command COMMAND, args ...any)
	Item() *NoteItem
}
