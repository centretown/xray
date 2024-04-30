package notes

import (
	"log"

	"unknwon.dev/i18n"
)

type Language struct {
	Code     string
	Title    string
	Base     []byte
	Source   []byte
	locale   *i18n.Locale
	fallback *i18n.Locale
}

func (lg *Language) String() string {
	return lg.Title
}

func (lg *Language) Locale() *i18n.Locale {
	return lg.locale
}

type Languages struct {
	Items          map[string]*Language
	List           []*Language
	FallbackCode   string
	FallbackTitle  string
	FallbackSource []byte
	fallback       *i18n.Locale
	store          *i18n.Store
}

func InitLanguages(lgs *Languages) {
	var err error
	lgs.store = i18n.NewStore()
	lgs.fallback, err = lgs.store.AddLocale(lgs.FallbackCode,
		lgs.FallbackTitle, lgs.FallbackSource)
	if err != nil {
		log.Fatal(err)
	}
}

func (lgs *Languages) AddSources(lang *Language) {
	var err error
	if lang.Base != nil {
		lang.locale, err = lgs.store.AddLocale(lang.Code,
			lang.Title, lang.Source, lang.Base)
	} else if lang.Source != nil {
		lang.locale, err = lgs.store.AddLocale(lang.Code,
			lang.Title, lang.Source)
	} else {
		lang.locale = lgs.fallback
	}
	if err != nil {
		log.Fatal(err)
	}
	lang.fallback = lgs.fallback
}
