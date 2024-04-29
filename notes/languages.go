package notes

import (
	"log"

	"unknwon.dev/i18n"
)

type Language struct {
	Code   string
	Title  string
	Source []byte
	locale *i18n.Locale
}

func (lg *Language) String() string {
	return lg.Title
}

func (lg *Language) Locale() *i18n.Locale {
	return lg.locale
}

type Languages struct {
	Source []byte
	Items  map[string]*Language
	List   []*Language
	store  *i18n.Store
}

func InitLanguages(lgs *Languages) {
	lgs.store = i18n.NewStore()
}

func (lgs *Languages) AddSources(lang *Language, sources []byte, other []byte) {
	var err error
	if other == nil {
		lang.locale, err = lgs.store.AddLocale(lang.Code,
			lang.Title, sources)
	} else {
		lang.locale, err = lgs.store.AddLocale(lang.Code,
			lang.Title, sources, other)
	}

	if err != nil {
		log.Fatal(err)
	}
}
