package text

type TextItem int

const (
	MajorUsage TextItem = iota
	MinorUsage
	KeyUsage
	InstallUsage
	QuickUsage
	LastTextItem
	FirstItem = 0
)

//go:generate stringer -type=TextItem
//go:generate gostringer -type=TextItem
