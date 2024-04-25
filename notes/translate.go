package notes

import (
	"log"

	"github.com/centretown/xray/locale"
	"unknwon.dev/i18n"
)

type LanguageItem struct {
	Code   string
	Title  string
	Locale *i18n.Locale
}

type LanguageList struct {
	Store *i18n.Store
	Items map[string]*LanguageItem
	List  []*LanguageItem
}

func NewLanguageList() *LanguageList {
	var (
		err       error
		loc       *i18n.Locale
		languages = &LanguageList{
			List: []*LanguageItem{
				{
					Code:  "en_US",
					Title: "English",
				},
				{
					Code:  "fr",
					Title: "Français",
				},
			},
			Items: make(map[string]*LanguageItem),
		}
	)

	languages.Store = i18n.NewStore()
	store := languages.Store
	for i := range languages.List {
		item := languages.List[i]
		if item.Code == "en_US" {
			loc, err = store.AddLocale(item.Code, item.Title,
				[]byte(locale.Locale_en_US),
				[]byte(locale.Locale_en_US_Values),
				[]byte(locale.Locale_en_US_Flags),
			)
		} else if item.Code == "fr" {
			loc, err = store.AddLocale(item.Code, item.Title,
				[]byte(locale.Locale_en_US_Values),
				[]byte(locale.Locale_en_US),
				[]byte(locale.Locale_en_US_Flags),
				[]byte(locale.Locale_fr),
				[]byte(locale.Locale_fr_Values),
				[]byte(locale.Locale_fr_Flags),
			)
		}

		if err != nil {
			log.Fatal(err)
		}
		item.Locale = loc
	}

	for _, lang := range languages.List {
		languages.Items[lang.Code] = lang
	}
	return languages
}
