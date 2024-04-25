package notes

import (
	"golang.org/x/exp/constraints"
	"unknwon.dev/i18n"
)

type Note struct {
	Label   string
	Value   string
	Values  []any
	Action  func(command int)
	Refresh func(values ...any)
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
}

func NewNotes(list []*Note, locale *i18n.Locale) (nts *Notes) {
	length := len(list)
	nts = &Notes{
		Length:  length,
		List:    list,
		Outputs: make([]*Output, length),
	}

	for i := range nts.List {
		nts.Outputs[i] = &Output{}
	}
	return nts
}

func (nts *Notes) Fetch(language *LanguageItem) {
	var (
		token  *Note
		output *Output
		i      int
		locale = language.Locale
	)

	for i, token = range nts.List {
		if token.Refresh != nil {
			token.Refresh(token.Values...)
		}
		output = nts.Outputs[i]
		output.Label = locale.Translate(token.Label)
		output.Value = locale.Translate(token.Value, token.Values...)
	}
}

func (nts *Notes) Draw(i int, draw func(i int, label, value string)) {
	if i < nts.Length {
		draw(i, nts.Outputs[i].Label, nts.Outputs[i].Value)
	}
}

func IncrementMore[T constraints.Integer | constraints.Float](command int, value *T) {
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

func Increment[T constraints.Integer | constraints.Float](command int, value *T) {
	switch command {
	case INC_MORE, INC:
		*value++
	case DEC_MORE, DEC:
		if *value-1 > 0 {
			*value--
		}
	}
}

func Select(command int, selection *int, length int) {
	switch command {
	case INC_MORE, INC:
		if *selection+1 >= length {
			*selection = 0
		} else {
			*selection++
		}
	case DEC_MORE, DEC:
		if *selection-1 >= 0 {
			*selection--
		} else {
			*selection = length - 1
		}
	}
}
