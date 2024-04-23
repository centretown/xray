package message

import (
	"fmt"
)

type Options struct {
	Sep      string
	TokenSep string
}

type Token struct {
	Label  Message
	Format string
	Values []any
}

type Output struct {
	Label string
	Value string
}

func Build(options *Options, tokens ...*Token) (outputs []*Output) {

	outputs = make([]*Output, len(tokens))
	var (
		token     *Token
		output    *Output
		i         int
		tokenLast = len(tokens) - 1
	)

	for i, token = range tokens {
		output = &Output{}
		output.Label = fmt.Sprintf(token.Label.String() + options.Sep)
		if i == tokenLast {
			output.Value = fmt.Sprintf(token.Format, token.Values...)
		} else {
			output.Value = fmt.Sprintf(token.Format+options.TokenSep, token.Values...)
		}
		output.Value = fmt.Sprintf(token.Format+options.TokenSep, token.Values...)
		outputs[i] = output
	}

	return outputs
}
