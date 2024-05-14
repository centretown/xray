package gizzmo

import (
	"github.com/centretown/xray/notes"

	rl "github.com/centretown/raylib-go/raylib"
)

func (gs *Game) BuildNotes() {
	var (
		content = &gs.Content
		options = content.options
	)
	options.FontSize = float64(content.Layout.Fontsize)
	// content.Languages = notes.NewLanguageList()
	// content.LanguageIndex = 1
	// content.Language = content.Languages.List[content.LanguageIndex]
	gs.RefreshEnvironment()
}

// from input thread (no raylib main thread action here)
func (gs *Game) updateState(command notes.Command) {
	var (
		content = &gs.Content
	)

	if command == notes.OPTIONS {
		content.commandState = !content.commandState
		return
	}

	content.options.Do(command)
	content.Layout.Refresh(int32(content.options.FontSize),
		content.options.Current)

}

// main thread only
func (gs *Game) RefreshEnvironment() {
	options := gs.Content.options

	num := rl.GetCurrentMonitor()
	mon := options.GetMonitor()
	mon.Num = num
	mon.Width = rl.GetMonitorWidth(num)
	mon.Height = rl.GetMonitorHeight(num)
	mon.RefreshRate = rl.GetMonitorRefreshRate(num)

	options.FrameRate = int64(rl.GetFPS())

	scr := options.GetScreen()
	scr.Width = int64(rl.GetScreenWidth())
	scr.Height = int64(rl.GetScreenHeight())
}
