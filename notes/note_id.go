package notes

// labels
const (
	Monitor   = "labels::monitor"
	View      = "labels::view"
	Duration  = "labels::duration"
	FrameRate = "labels::framerate"
	Capture   = "labels::capture"
	Language  = "labels::language"
	FontSize  = "labels::fontsize"
)

// values
const (
	MonitorValue   = "values::monitor"
	ViewValue      = "values::view"
	DurationValue  = "values::duration"
	FrameRateValue = "values::framerate"
	CaptureValue   = "values::capture"
	LanguageValue  = "values::language"
	StringValue    = "values::string"
	IntegerValue   = "values::integer"
	FloatValue     = "values::float"
)

// flags
const (
	Major        = "options::major"
	MajorShort   = "options::majorshort"
	MajorUsage   = "options::majorusage"
	Minor        = "options::minor"
	MinorShort   = "options::minorshort"
	MinorUsage   = "options::minorusage"
	Key          = "options::key"
	KeyShort     = "options::keyshort"
	KeyUsage     = "options::keyusage"
	Install      = "options::install"
	InstallShort = "options::installshort"
	InstallUsage = "options::installusage"
	Quick        = "options::quick"
	QuickShort   = "options::quickshort"
	QuickUsage   = "options::quickusage"
)

const (
	NONE = iota
	HELP
	MORE
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
