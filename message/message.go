package message

import (
	"log"

	"github.com/centretown/xray/locale"

	"unknwon.dev/i18n"
)

var Locale *i18n.Locale

func init() {
	var err error
	s := i18n.NewStore()
	Locale, err = s.AddLocale("en-US", "English",
		[]byte(locale.Locale_en_US_Values),
		[]byte(locale.Locale_en_US),
		[]byte(locale.Locale_fr_Values),
		[]byte(locale.Locale_fr),
	)
	if err != nil {
		log.Fatal(err)
	}
}
