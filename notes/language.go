package notes

import "unknwon.dev/i18n"

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
