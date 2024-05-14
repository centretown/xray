package notebooks

import "github.com/centretown/xray/notes"

type CaptureBook struct {
	Notebook *notes.Notebook
}

func NewCaptureBook(languageChooser *notes.LanguageChooser) *CaptureBook {
	cpbk := &CaptureBook{
		Notebook: notes.NewNotebook(languageChooser),
	}
	cpbk.Notebook.Add(
		NewCaptureEntry(),
	)
	return cpbk
}
