package notes

import "golang.org/x/exp/constraints"

type Note struct {
	Label     string
	Value     string
	Values    []any
	Action    func(command int)
	UpdateAll func(values ...any)
}

func (nt *Note) CanAct() bool {
	return nt.Action != nil
}

type Output struct {
	Label string
	Value string
}

type Notes struct {
	List    []*Note
	Outputs []*Output
	Length  int
	Current int
}

func NewNotes(tokens []*Note) (tkl *Notes) {
	length := len(tokens)
	tkl = &Notes{
		Length:  length,
		List:    tokens,
		Outputs: make([]*Output, length),
	}

	for i := range tkl.List {
		tkl.Outputs[i] = &Output{}
	}
	return tkl
}

func (tkl *Notes) Fetch() {
	var (
		token  *Note
		output *Output
		i      int
	)

	for i, token = range tkl.List {
		if token.UpdateAll != nil {
			token.UpdateAll(token.Values...)
		}
		output = tkl.Outputs[i]
		output.Label = Current.Translate(token.Label)
		output.Value = Current.Translate(token.Value, token.Values...)
	}
}

func (tkl *Notes) Draw(i int, draw func(i int, label, value string)) {
	if i < tkl.Length {
		draw(i, tkl.Outputs[i].Label, tkl.Outputs[i].Value)
	}
}

func Increment[T constraints.Integer | constraints.Float](command int, value *T) {
	switch command {
	case INC_MORE:
		*value += 10
	case INC:
		*value++
	case DEC_MORE:
		if *value-10 > 0 {
			*value -= 10
		} else {
			*value = 1
		}
	case DEC:
		if *value-1 > 0 {
			*value--
		}
	}
}
