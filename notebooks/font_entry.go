package notebooks

import "github.com/centretown/xray/notes"

type FontEntry struct {
	*notes.Ranger[float64]
}

func NewFontEntry(fontsize *float64) *FontEntry {
	fe := &FontEntry{}
	fe.Ranger = notes.NewRanger(notes.FontsizeLabel, notes.FloatValue,
		fontsize, 8, 100, 10)
	var _ notes.Note = fe
	return fe
}

func (fe *FontEntry) Values() []any {
	return []any{*fe.Ranger.Current}
}
