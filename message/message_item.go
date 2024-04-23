package message

type MessageItem int

const (
	MajorUsage MessageItem = iota
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
	FPS
	LastTextItem
	FirstItem = 0
)

//go:generate stringer -type=MessageItem
