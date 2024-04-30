package builder

import (
	"github.com/centretown/xray/notes"
)

func InitLanguageList(languages *notes.Languages) {

	notes.InitLanguages(languages)

	for _, lang := range languages.List {
		languages.AddSources(lang)
	}

	for _, lang := range languages.List {
		languages.Items[lang.Code] = lang
	}

}
