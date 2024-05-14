package notebooks

import "github.com/centretown/xray/notes"

type Screen struct {
	Width  int64
	Height int64
}

type ScreenEntry struct {
	*notes.Ranger[int]
	Screens []*Screen
	Current int
}

func NewScreenEntry(screens ...*Screen) *ScreenEntry {
	scr := &ScreenEntry{
		Screens: make([]*Screen, 0),
	}
	scr.Screens = append(scr.Screens, screens...)
	scr.Ranger = notes.NewRanger(notes.ViewLabel, notes.ViewValue,
		&scr.Current, 0, len(screens), 1)
	scr.CanDo = true
	var _ notes.Note = scr
	return scr
}

// func (scr *ScreenEntry) Do(command notes.Command, args ...any) {}

func (scr *ScreenEntry) Values() []any {
	screen := scr.Screens[scr.Current]
	return []any{
		screen.Width,
		screen.Height,
	}
}

func (scr *ScreenEntry) Add(screen *Screen) {
	scr.Screens = append(scr.Screens, screen)
	scr.Ranger.Max = len(scr.Screens)
}

func (scr *ScreenEntry) Get() (screen *Screen) {
	return scr.Screens[scr.Current]
}
