package notebooks

import "github.com/centretown/xray/notes"

type Monitor struct {
	Num         int
	Width       int
	Height      int
	RefreshRate int
}

type MonitorEntry struct {
	*notes.Ranger[int]
	Monitors []*Monitor
	Current  int
}

func NewMonitorEntry(monitors ...*Monitor) *MonitorEntry {
	mon := &MonitorEntry{
		Monitors: make([]*Monitor, 0),
	}
	mon.Monitors = append(mon.Monitors, monitors...)
	mon.Ranger = notes.NewRanger(notes.MonitorLabel, notes.MonitorValue,
		&mon.Current, 0, len(monitors), 1)

	mon.CanDo = true
	var _ notes.Note = mon
	return mon
}

func (mon *MonitorEntry) Values() []any {
	current := mon.Monitors[mon.Current]
	return []any{
		current.Num,
		current.Width,
		current.Height,
		current.RefreshRate,
	}
}

func (mon *MonitorEntry) Add(monitor *Monitor) {
	mon.Monitors = append(mon.Monitors, monitor)
	mon.Ranger.Max = len(mon.Monitors)
}

func (mon *MonitorEntry) Get() (monitor *Monitor) {
	return mon.Monitors[mon.Current]
}
