package notes

type Scribble struct {
	Label string
	Value string
}

type Scribe struct {
	LabelKey  string
	FormatKey string
	Output    Scribble
	CanDo     bool
}

type Note interface {
	Values() []any
	Do(command Command, args ...any)
	GetScribe() *Scribe
}
