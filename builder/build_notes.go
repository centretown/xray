package builder

import (
	"github.com/centretown/xray/builder/locale"
	"github.com/centretown/xray/entries"
	"github.com/centretown/xray/gizzmo"
	"github.com/centretown/xray/notes"
)

func BuildNotes(gs *gizzmo.Game) {
	content := &gs.Content
	lang := &content.Languages
	lang.List = []*notes.Language{
		{
			Code:   "en_US",
			Title:  "English",
			Source: nil,
		},
		{
			Code:   "fr",
			Title:  "Français",
			Source: []byte(locale.Locale_fr),
		},
	}
	lang.Source = []byte(locale.Locale_en_US)
	lang.Items = make(map[string]*notes.Language)

	InitLanguageList(&content.Languages)
	content.LanguageIndex = 1
	content.Language = content.Languages.List[content.LanguageIndex]

	content.CaptureNotes = notes.NewNotes()
	buildOptionsNotes(gs)
	buildCaptureNotes(gs)
	buildKeyNotes(gs)
	buildPadNotes(gs)
}

const (
	NONE = iota
	LANGUAGE
	FONTSIZE
)

func buildOptionsNotes(gs *gizzmo.Game) {
	content := &gs.Content
	content.CaptureNotes.Add(
		entries.NewLanguageEntry(&content.Languages),
		entries.NewFontEntry(&content.FontSize),
		entries.NewMonitorEntry(&content.Monitor),
		entries.NewScreenView(&content.Screen))
}

func buildCaptureNotes(gs *gizzmo.Game) {
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
