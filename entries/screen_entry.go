package entries

import "github.com/centretown/xray/notes"

type Screen struct {
	Width  int64
	Height int64
}

type ScreenEntry struct {
	notes.Composite[Screen]
}

func NewScreenEntry(view *Screen) *ScreenEntry {
	scr := &ScreenEntry{}
	notes.SetupComposite(&scr.Composite, notes.ViewLabel, notes.ViewValue, view)
	var _ notes.Note = scr
	return scr
}

func (scr *ScreenEntry) Do(command notes.COMMAND, args ...any) {}

func (scr *ScreenEntry) Values() []any {
	custom := scr.Custom
	return []any{
		custom.Width,
		custom.Height,
	}
}
