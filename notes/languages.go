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

func (lgs *Languages) Store() *i18n.Store {
	return lgs.store
}

func (lgs *Languages) Add(lang *Language, sources []byte, other []byte) {
	var err error
	if other == nil {
		lang.locale, err = lgs.store.AddLocale(lang.Code, lang.Title,
			sources)
	} else {
		lang.locale, err = lgs.store.AddLocale(lang.Code, lang.Title,
			sources, other)
	}

	if err != nil {
		log.Fatal(err)
	}
}

// func NewLanguageList() *Languages {
// 	var (
// 		err       error
// 		loc       *i18n.Locale
// 		languages = &Languages{
// 			List: []*Language{
// 				{
// 					Code:  "en_US",
// 					Title: "English",
// 				},
// 				{
// 					Code:   "fr",
// 					Title:  "Français",
// 					Source: []byte(locale.Locale_fr),
// 				},
// 			},
// 			Source: []byte(locale.Locale_en_US),
// 			Items:  make(map[string]*Language),
// 		}
// 	)

// 	languages.store = i18n.NewStore()
// 	store := languages.store
// 	for i := range languages.List {
// 		item := languages.List[i]
// 		if item.Code == "en_US" {
// 			loc, err = store.AddLocale(item.Code, item.Title,
// 				languages.Source,
// 			)
// 		} else if item.Code == "fr" {
// 			loc, err = store.AddLocale(item.Code, item.Title,
// 				languages.Source,
// 				item.Source,
// 			)
// 		}

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		item.locale = loc
// 	}

// 	for _, lang := range languages.List {
// 		languages.Items[lang.Code] = lang
// 	}
// 	return languages
// }
