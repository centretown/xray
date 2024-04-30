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

func BuildGameNotes(gs *gizzmo.Game) {

	content := &gs.Content
	languages := &content.Languages
	languages.FallbackCode = "en_US"
	languages.FallbackTitle = "English (US)"
	languages.FallbackSource = []byte(locale.Locale_en_US)

	languages.List = []*notes.Language{
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
	languages.Items = make(map[string]*notes.Language)

	InitLanguageList(&content.Languages)
	chooser := notes.NewLanguageChooser(languages)
	content.Language = content.Languages.List[content.LanguageCurrent]

	content.Options = notes.NewNotes(chooser)
	content.Options.Add(
		chooser,
		entries.NewFontEntry(&content.Fontsize),
		entries.NewMonitorEntry(&content.Monitor),
		entries.NewScreenEntry(&content.Screen))

	content.Capture = notes.NewNotes(chooser)
	content.Capture.Add()

	buildKeyNotes(gs)
	buildPadNotes(gs)
}

const (
	NONE = iota
	LANGUAGE
	FOntsIZE
)

func buildCaptureNotes(gs *gizzmo.Game, chooser *notes.LanguageChooser) {
	// content := &gs.Content
	// list := make([]*notes.Note, 0)
	// list = append(list,
	// 	&notes.Note{Label: notes.Capture, Value: notes.CaptureValue,
	// 		Values: []any{content.captureTotal, content.captureCount},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.captureTotal
	// 			values[1] = content.captureCount
	// 		}})
	// content.CaptureNotes = notes.NewNotes(list,
	// 	content.Languages.List[content.LanguageIndex].Locale)
}

func buildKeyNotes(gs *gizzmo.Game) {
}

func buildPadNotes(gs *gizzmo.Game) {
}
