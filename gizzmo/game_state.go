package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
)

func (gs *Game) BuildNotes() {
	// var (
	// 	content = &gs.Content
	// )
	// content.Languages = notes.NewLanguageList()
	// content.LanguageIndex = 1
	// content.Language = content.Languages.List[content.LanguageIndex]
	// gs.RefreshEnvironment()
}

// from input thread (no raylib main thread action here)
func (gs *Game) updateState(command notes.COMMAND) {
	var (
		content = &gs.Content
		length  = content.Options.Length
		// note    *notes.Note
	)

	switch command {
	case notes.OPTIONS:
		content.commandState = !content.commandState

	case notes.NEXT:
		if content.OptionCurrent+1 < length {
			content.OptionCurrent++
		} else {
			content.OptionCurrent = 0
		}
		content.Layout.Current = content.OptionCurrent

	case notes.PREVIOUS:
		if content.OptionCurrent-1 >= 0 {
			content.OptionCurrent--
		} else {
			content.OptionCurrent = length - 1
		}
		content.Layout.Current = content.OptionCurrent

	// case notes.INCREMENT, notes.INCREMENT_MORE, notes.DECREMENT, notes.DECREMENT_MORE:
	// 	note = content.OptionsNotes.Notes[content.CurrentOption]
	// 	if note.CanAct() {
	// 		note.Act(command)
	// 	}
	case notes.SHARE:
		if content.capturing {
			content.endCapturing = true
		} else {
			content.beginCapturing = true
		}
	case notes.PAUSE_PLAY:
		content.paused = !content.paused
	}
}

// main thread only
func (gs *Game) RefreshEnvironment() {
	content := &gs.Content
	mon := &content.Monitor
	mon.Num = rl.GetCurrentMonitor()
	mon.Width = rl.GetMonitorWidth(mon.Num)
	mon.Height = rl.GetMonitorHeight(mon.Num)
	mon.RefreshRate = rl.GetMonitorRefreshRate(mon.Num)
	content.CurrentFrameRate = int64(rl.GetFPS())
	scr := &content.Screen
	scr.Width = int64(rl.GetScreenWidth())
	scr.Height = int64(rl.GetScreenHeight())
}
