package notebooks

import "github.com/centretown/xray/notes"

type CaptureEntry struct {
	notes.Scribe
	Count int64
	Total int64
}

func NewCaptureEntry() *CaptureEntry {
	cap := &CaptureEntry{
		Scribe: notes.Scribe{
			LabelKey:  notes.CaptureLabel,
			FormatKey: notes.CaptureValue,
		},
	}
	var _ notes.Note = cap
	return cap
}
func (cap *CaptureEntry) GetScribe() *notes.Scribe {
	return &cap.Scribe
}

func (cap *CaptureEntry) Do(command notes.Command, args ...any) {
	switch command {
	case notes.SET:
		{
			var (
				total int64
				ok    bool
			)
			if len(args) > 0 {
				total, ok = args[0].(int64)
				if ok {
					cap.Total = total
				}
			}
		}

	case notes.INCREMENT:
		cap.Count++
	}
}

func (scr *CaptureEntry) Values() []any {
	return []any{
		scr.Count,
		scr.Total,
	}
}
