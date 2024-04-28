package notes

type MonitorItem struct {
	Num         int
	Width       int
	Height      int
	RefreshRate int
}

type Monitor struct {
	Composite[MonitorItem]
}

func NewMonitor(label, format string, monitor *MonitorItem) *Monitor {
	mon := &Monitor{}
	InitComposite(&mon.Composite, label, format, monitor)
	var _ Note = mon
	return mon
}

func (mon *Monitor) Do(command COMMAND) {}

func (mon *Monitor) Values() []any {
	custom := mon.Custom
	return []any{
		custom.Num,
		custom.Width,
		custom.Height,
		custom.RefreshRate,
	}
}
