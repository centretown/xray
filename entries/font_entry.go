package entries

import "github.com/centretown/xray/notes"

func NewFontEntry(fontsize *float64) *notes.Ranger[float64] {
	return notes.NewRanger(notes.FontsizeLabel, notes.FloatValue,
		fontsize, 8, 100, 10)
}
