package builder

import (
	"github.com/centretown/xray/notes"
)

func SetupVocabulary(vocabulary *notes.VocabularyItem) {

	notes.SetupVocabulary(vocabulary)

	for _, lang := range vocabulary.Languages {
		vocabulary.AddSources(lang)
	}
}
