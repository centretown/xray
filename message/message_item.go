package message

type Message int

const (
	Monitor        = "messages::monitor"
	MonitorValue   = "values::monitor"
	View           = "messages::view"
	ViewValue      = "values::view"
	Duration       = "messages::duration"
	DurationValue  = "values::duration"
	FrameRate      = "messages::framerate"
	FrameRateValue = "values::framerate"
	Capture        = "messages::capture"
	CaptureValue   = "values::capture"
	LastTextItem
	FirstItem = 0
)
