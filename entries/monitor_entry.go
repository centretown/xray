package entries

import "github.com/centretown/xray/notes"

type Monitor struct {
	Num         int
	Width       int
	Height      int
	RefreshRate int
}

type MonitorEntry struct {
	notes.Composite[Monitor]
}

func NewMonitorEntry(monitor *Monitor) *MonitorEntry {
	mon := &MonitorEntry{}
	notes.InitComposite(&mon.Composite,
		notes.MonitorLabel, notes.MonitorValue,
		monitor)
	var _ notes.Note = mon
	return mon
}

func (mon *MonitorEntry) Do(command notes.COMMAND, args ...any) {}

func (mon *MonitorEntry) Values() []any {
	custom := mon.Custom
	return []any{
		custom.Num,
		custom.Width,
		custom.Height,
		custom.RefreshRate,
	}
}
