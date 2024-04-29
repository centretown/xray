package entries

import "github.com/centretown/xray/notes"

func NewLanguageEntry(languages *notes.Languages) *notes.Chooser[*notes.Language] {
	return notes.NewChooser(notes.LanguageLabel, notes.StringValue,
		&languages.List)
}
