package builder

import (
	"github.com/centretown/xray/builder/locale"
	"github.com/centretown/xray/notes"
)

func NewLanguageList() *notes.Languages {
	var (
		languages = &notes.Languages{
			List: []*notes.Language{
				{
					Code:  "en_US",
					Title: "English",
				},
				{
					Code:   "fr",
					Title:  "Français",
					Source: []byte(locale.Locale_fr),
				},
			},
			Source: []byte(locale.Locale_en_US),
			Items:  make(map[string]*notes.Language),
		}
	)

	notes.InitLanguages(languages)

	for i := range languages.List {
		item := languages.List[i]
		languages.Add(item, languages.Source, item.Source)
	}

	for _, lang := range languages.List {
		languages.Items[lang.Code] = lang
	}
	return languages
}
