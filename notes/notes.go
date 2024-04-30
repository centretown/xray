package notes

type Doer interface {
	Values() []any
	// Bind(value any)
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

type Notebook struct {
	Language *LanguageChooser
	Notes    []Note
	Length   int
}

func NewNotes(languageChooser *LanguageChooser) (ntbk *Notebook) {
	return &Notebook{
		Notes:    make([]Note, 0),
		Language: languageChooser,
	}
}

func (ntbk *Notebook) Add(notes ...Note) {
	ntbk.Notes = append(ntbk.Notes, notes...)
}

func (ntbk *Notebook) Fetch() {
	var (
		language *Language = ntbk.Language.Current()
		locale             = language.locale
		note     Note
		output   *Output
		item     *NoteItem
	)

	for _, note = range ntbk.Notes {
		item = note.Item()
		output = &item.Output
		output.Label = locale.TranslateWithFallback(language.fallback, item.LabelKey)
		output.Value = locale.TranslateWithFallback(language.fallback, item.FormatKey, note.Values()...)
	}
}

func (ntbk *Notebook) Draw(i int, draw func(i int, label, value string)) {
	if i < ntbk.Length {
		output := &ntbk.Notes[i].Item().Output
		draw(i, output.Label, output.Value)
	}
}

func (ntbk *Notebook) DrawAll(draw func(i int, label, value string)) {
	for i := range ntbk.Notes {
		output := &ntbk.Notes[i].Item().Output
		draw(i, output.Label, output.Value)
	}
}
