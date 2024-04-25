package notes

// labels
const (
	Monitor   = "labels::monitor"
	View      = "labels::view"
	Duration  = "labels::duration"
	FrameRate = "labels::framerate"
	Capture   = "labels::capture"
)

// values
const (
	MonitorValue   = "values::monitor"
	ViewValue      = "values::view"
	DurationValue  = "values::duration"
	FrameRateValue = "values::framerate"
	CaptureValue   = "values::capture"
)

// flags
const (
	MajorFlag    = "options::major"
	MajorShort   = "options::majorshort"
	MajorUsage   = "options::majorusage"
	MinorFlag    = "options::minor"
	MinorShort   = "options::minorshort"
	MinorUsage   = "options::minorusage"
	KeyFlag      = "options::key"
	KeyShort     = "options::keyshort"
	KeyUsage     = "options::keyusage"
	InstallFlag  = "options::install"
	InstallShort = "options::installshort"
	InstallUsage = "options::installusage"
	QuickFlag    = "options::quick"
	QuickShort   = "options::quickshort"
	QuickUsage   = "options::quickusage"
)

const (
	MORE = iota
	NEXT_NOTE
	PREV_NOTE
	INC
	DEC
	INC_MORE
	DEC_MORE
	PAUSE_PLAY
	CAPTURE
	COMMANDS
)
