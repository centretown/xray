package notes

type Doer interface {
	Values() []any
	Bind(value any)
	Do(command COMMAND)
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

type NoteG[T Doer] struct {
	Item NoteItem
	Doer T
}

type Note interface {
	Doer
	Item() *NoteItem
}

type Notes struct {
	Notes  []Note
	Length int
}

func NewNotes() (nts *Notes) {
	return &Notes{
		Notes: make([]Note, 0),
	}
}

func (nts *Notes) Add(note Note) {
	nts.Notes = append(nts.Notes, note)
}

func (nts *Notes) Fetch(language *Language) {
	var (
		locale = language.locale
		note   Note
		output *Output
		item   *NoteItem
	)

	for _, note = range nts.Notes {
		item = note.Item()
		output = &item.Output
		output.Label = locale.Translate(item.LabelKey)
		output.Value = locale.Translate(item.FormatKey, note.Values()...)
	}
}

func (nts *Notes) Draw(i int, draw func(i int, label, value string)) {
	if i < nts.Length {
		output := &nts.Notes[i].Item().Output
		draw(i, output.Label, output.Value)
	}
}

func (nts *Notes) DrawAll(draw func(i int, label, value string)) {
	for i := range nts.Notes {
		output := &nts.Notes[i].Item().Output
		draw(i, output.Label, output.Value)
	}
}
