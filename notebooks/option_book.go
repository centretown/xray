package notebooks

import (
	"fmt"

	"github.com/centretown/xray/notes"
)

type OptionBook struct {
	Notebook  *notes.Notebook
	Monitor   Monitor
	Screen    Screen
	FontSize  float64
	FrameRate int64

	LanguageChooser *notes.LanguageChooser
	FontEntry       *FontEntry
	MonitorEntry    *MonitorEntry
	ScreenEntry     *ScreenEntry

	Current int32
}

func NewOptionBook(languageChooser *notes.LanguageChooser) *OptionBook {
	opbk := &OptionBook{
		Notebook:        notes.NewNotebook(languageChooser),
		LanguageChooser: languageChooser,
	}

	opbk.FontEntry = NewFontEntry(&opbk.FontSize)
	opbk.MonitorEntry = NewMonitorEntry(&opbk.Monitor)
	opbk.ScreenEntry = NewScreenEntry(&opbk.Screen)

	opbk.Notebook.Add(
		languageChooser,
		opbk.FontEntry,
		opbk.MonitorEntry,
		opbk.ScreenEntry,
	)

	return opbk
}

func (opbk *OptionBook) GetMonitor() *Monitor {
	return &opbk.Monitor
}

func (opbk *OptionBook) MonitorSet(num, width, height, refreshRate int) {
	mon := opbk.Monitor
	mon.Num = num
	mon.Width = width
	mon.Height = height
	mon.RefreshRate = refreshRate

}

func (opbk *OptionBook) GetScreen() *Screen {
	return &opbk.Screen
}

func (opbk *OptionBook) Do(command notes.Command) (current int32) {
	current = opbk.Current
	fmt.Println(current, opbk.Notebook.Length, command.String())

	switch command {
	case notes.NEXT:
		current++
	case notes.PREVIOUS:
		current--
	default:
		// case notes.INCREMENT, notes.INCREMENT_MORE, notes.DECREMENT, notes.DECREMENT_MORE:
		opbk.Notebook.Notes[current].Do(command)
		return
	}

	if current < 0 {
		current = opbk.Notebook.Length - 1
	} else if current >= opbk.Notebook.Length {
		current = 0
	}
	opbk.Current = current
	return
}
