package notes

type Notebook struct {
	Language *LanguageChooser
	Notes    []Note
	Length   int32
	Current  int32
}

func NewNotebook(languageChooser *LanguageChooser) (ntbk *Notebook) {
	return &Notebook{
		Notes:    make([]Note, 0),
		Language: languageChooser,
	}
}

func (ntbk *Notebook) Add(notes ...Note) {
	ntbk.Notes = append(ntbk.Notes, notes...)
	ntbk.Length = int32(len(ntbk.Notes))
}

func (ntbk *Notebook) Fetch() {
	var (
		language *Language = ntbk.Language.Current()
		locale             = language.locale
		note     Note
		output   *Scribble
		scribe   *Scribe
	)

	for _, note = range ntbk.Notes {
		scribe = note.GetScribe()
		output = &scribe.Output
		output.Label = locale.TranslateWithFallback(language.fallback, scribe.LabelKey)
		output.Value = locale.TranslateWithFallback(language.fallback, scribe.FormatKey,
			note.Values()...)
	}
}

func (ntbk *Notebook) Draw(i int, draw func(i int, label, value string)) {
	if i < int(ntbk.Length) {
		output := &ntbk.Notes[i].GetScribe().Output
		draw(i, output.Label, output.Value)
	}
}

func (ntbk *Notebook) DrawAll(draw func(i int, label, value string)) {
	for i := range ntbk.Notes {
		output := &ntbk.Notes[i].GetScribe().Output
		draw(i, output.Label, output.Value)
	}
}
