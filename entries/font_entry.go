package entries

import "github.com/centretown/xray/notes"

func NewFontEntry(fontSize *float64) *notes.Ranger[float64] {
	return notes.NewRanger(notes.FontSizeLabel, notes.FloatValue,
		fontSize, 8, 100, 10)
}
