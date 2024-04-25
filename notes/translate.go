package notes

import (
	"log"

	"github.com/centretown/xray/locale"

	"unknwon.dev/i18n"
)

var Current *i18n.Locale

func Initialize() {
	var err error
	s := i18n.NewStore()
	Current, err = s.AddLocale("en-US", "English",
		[]byte(locale.Locale_en_US_Values),
		[]byte(locale.Locale_en_US),
		[]byte(locale.Locale_en_US_Flags),
		[]byte(locale.Locale_fr_Values),
		[]byte(locale.Locale_fr),
		[]byte(locale.Locale_fr_Flags),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Translate(key string, args ...interface{}) string {
	return Current.Translate(key, args)
}
