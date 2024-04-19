package en

import "github.com/centretown/xray/text"

var items = []string{
	"major version number",    // majorUsage
	"minor version number",    // minorUsage
	"uuid key",                // keyUsage
	"install build to folder", // installUsage
	"quick build and run in temporary memory database", // quickUsage
}

// var _ text.Texter = (*En)(nil)

type En struct {
	Items [text.LastTextItem]string
}

func (en *En) Format(item text.TextItem) string {
	return items[int(item)]
}