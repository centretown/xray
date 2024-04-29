package builder

import (
	"github.com/centretown/xray/notes"
)

func InitLanguageList(languages *notes.Languages) *notes.Languages {

	notes.InitLanguages(languages)

	for i := range languages.List {
		item := languages.List[i]
		languages.AddSources(item, languages.Source, item.Source)
	}

	for _, lang := range languages.List {
		languages.Items[lang.Code] = lang
	}

	return languages
}
