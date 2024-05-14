package notes

import (
	"log"

	"github.com/centretown/xray/class"
	"github.com/centretown/xray/gizzmodb/model"
	"unknwon.dev/i18n"
)

type VocabularyItem struct {
	// LanguageMap    map[string]*Language
	Languages      []*Language
	FallbackCode   string
	FallbackTitle  string
	FallbackSource []byte
	fallback       *i18n.Locale
	store          *i18n.Store
}

type Vocabulary struct {
	model.RecorderClass[VocabularyItem]
}

func NewVocabulary() {
	voc := &Vocabulary{}
	model.SetupRecorder[VocabularyItem](voc, class.Vocabulary.String(), int32(class.Vocabulary))
}

func SetupVocabulary(voc *VocabularyItem) {
	var err error
	voc.store = i18n.NewStore()
	voc.fallback, err = voc.store.AddLocale(voc.FallbackCode,
		voc.FallbackTitle, voc.FallbackSource)
	if err != nil {
		log.Fatal(err)
	}
}

func (voc *VocabularyItem) AddSources(lang *Language) {
	var err error
	if lang.Base != nil {
		lang.locale, err = voc.store.AddLocale(lang.Code,
			lang.Title, lang.Source, lang.Base)
	} else if lang.Source != nil {
		lang.locale, err = voc.store.AddLocale(lang.Code,
			lang.Title, lang.Source)
	} else {
		lang.locale = voc.fallback
	}
	if err != nil {
		log.Fatal(err)
	}
	lang.fallback = voc.fallback
}
