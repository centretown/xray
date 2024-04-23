package message

import (
	"fmt"
	"strings"
)

type Options struct {
	Sep      string
	TokenSep string
}

type Token struct {
	Item   MessageItem
	Format string
	Values []any
}

func Message(options *Options, tokens ...*Token) string {
	var (
		bld   = strings.Builder{}
		token *Token
	)

	for _, token = range tokens {
		bld.WriteString(fmt.Sprintf(
			token.Item.String()+
				options.Sep+
				token.Format+
				options.TokenSep, token.Values...))
	}

	return bld.String()
}
