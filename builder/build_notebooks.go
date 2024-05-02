package builder

import (
	"github.com/centretown/xray/builder/locale"
	"github.com/centretown/xray/entries"
	"github.com/centretown/xray/gizzmo"
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
	vocabulary.LanguageMap = make(map[string]*notes.Language)

	SetupVocabulary(vocabulary)

	chooser := notes.NewLanguageChooser(vocabulary)
	// content.Language = content.Languages.List[content.LanguageCurrent]

	var (
		monitor     entries.Monitor
		screen      entries.Screen
		fontsize    float64
		capture     entries.Capture
		optionBook  *notes.Notebook
		captureBook *notes.Notebook
	)

	optionBook = notes.NewNotebook(chooser)
	optionBook.Add(
		chooser,
		entries.NewFontEntry(&fontsize),
		entries.NewMonitorEntry(&monitor),
		entries.NewScreenEntry(&screen))

	captureBook = notes.NewNotebook(chooser)
	captureBook.Add(
		entries.NewCaptureEntry(&capture),
	)
	game.SetOptions(optionBook, captureBook)
}
