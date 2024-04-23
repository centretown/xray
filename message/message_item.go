package message

type Message int

const (
	MajorUsage Message = iota
	MinorUsage
	KeyUsage
	InstallUsage
	QuickUsage
	Monitor
	View
	Capture
	Duration
	Frames
	Capturing
	Mhz
	Counter
	FrameRate
	Captured
	Current
	LastTextItem
	FirstItem = 0
)

//go:generate stringer -type=Message
