package message

type Token struct {
	Label  string
	Value  string
	Values []any
	Update func(values ...any)
}

type Output struct {
	Label string
	Value string
}

type TokenList struct {
	List    []*Token
	Outputs []*Output
	Length  int
	Current int
}

func NewTokenList(tokens []*Token) (tkl *TokenList) {
	length := len(tokens)
	tkl = &TokenList{
		Length:  length,
		List:    tokens,
		Outputs: make([]*Output, length),
	}

	for i := range tkl.List {
		tkl.Outputs[i] = &Output{}
	}
	return tkl
}

func (tkl *TokenList) Fetch() {
	var (
		token  *Token
		output *Output
		i      int
	)

	for i, token = range tkl.List {
		if token.Update != nil {
			token.Update(token.Values...)
		}
		output = tkl.Outputs[i]
		output.Label = Locale.Translate(token.Label)
		output.Value = Locale.Translate(token.Value, token.Values...)
	}
}

// DrawText(text string, posX int32, posY int32, fontSize int32, col color.RGBA)

func (tkl *TokenList) Draw(i int, draw func(i int, label, value string)) {
	if i < tkl.Length {
		draw(i, tkl.Outputs[i].Label, tkl.Outputs[i].Value)
	}
}
