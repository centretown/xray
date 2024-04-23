package message

import (
	"log"

	"github.com/centretown/xray/message/locale"

	"unknwon.dev/i18n"
)

type Token struct {
	Label  string
	Value  string
	Values []any
}

type Output struct {
	Label string
	Value string
}

var English *i18n.Locale

func init() {
	var err error
	s := i18n.NewStore()
	English, err = s.AddLocale("en-US", "English",
		[]byte(locale.Locale_en_US_Values),
		[]byte(locale.Locale_en_US),
		[]byte(locale.Locale_fr_Values),
		[]byte(locale.Locale_fr),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func Build(tokens ...*Token) (outputs []*Output) {
	outputs = make([]*Output, len(tokens))
	var (
		token  *Token
		output *Output
		i      int
	)

	for i, token = range tokens {
		output = &Output{}
		output.Label = English.Translate(token.Label)
		output.Value = English.Translate(token.Value, token.Values...)
		outputs[i] = output
	}

	return outputs
}
