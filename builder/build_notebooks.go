package builder

import (
	"github.com/centretown/xray/builder/locale"
	"github.com/centretown/xray/gizzmo"
	"github.com/centretown/xray/notebooks"
	"github.com/centretown/xray/notes"
)

const (
	LANG_BASE     = 0
	LANG_ENG_BASE = 0
	LANG_FR_BASE  = 1
)

var Vocabulary notes.VocabularyItem

func BuildNotebooks(game *gizzmo.Game) {

	vocabulary := &Vocabulary
	vocabulary.FallbackCode = "en_US"
	vocabulary.FallbackTitle = "English (US)"
	vocabulary.FallbackSource = []byte(locale.Locale_en_US)

	vocabulary.Languages = []*notes.Language{
		{
			Code:  "en_US",
			Title: "English (US)",
		},
		{
			Code:   "fr",
			Title:  "Français",
			Source: []byte(locale.Locale_fr),
		},
		{
			Code:   "fr_CA",
			Title:  "Français (CA)",
			Base:   []byte(locale.Locale_fr),
			Source: []byte(locale.Locale_fr_CA),
		},
		{
			Code:   "en_CA",
			Title:  "English (CA)",
			Source: []byte(locale.Locale_en_CA),
		},
	}

	SetupVocabulary(vocabulary)

	var (
		chooser     = notes.NewLanguageChooser(vocabulary)
		optionBook  = notebooks.NewOptionBook(chooser)
		captureBook = notebooks.NewCaptureBook(chooser)
	)
	game.AddNotebooks(optionBook, captureBook)
}
