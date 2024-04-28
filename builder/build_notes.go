package builder

import (
	"github.com/centretown/xray/gizzmo"
)

func BuildNotes(gs *gizzmo.Game) {
	content := &gs.Content
	content.Languages = NewLanguageList()
	content.LanguageIndex = 1
	content.Language = content.Languages.List[content.LanguageIndex]

	buildOptionsNotes(gs)
	buildCaptureNotes(gs)
	buildKeyNotes(gs)
	buildPadNotes(gs)
}

const (
	NONE = iota
	LANGUAGE
	FONTSIZE
)

func buildOptionsNotes(gs *gizzmo.Game) {
	// content := &gs.Content
	// nts := notes.NewNotes()

	// nts.Add(&notes.Note{
	// 	Id:    0,
	// 	Label: notes.LanguageLabel, Value: notes.StringValue,
	// 	Command: notes.CmdChoose,
	// 	Values:  []any{&content.Language.Title, len(content.Languages.List)},
	// })

	// content.OptionsNotes = nts

	// list = append(list,
	// 	&notes.Note{Label: notes.Language, Value: notes.StringValue,
	// 		Values: []any{content.Language.Title},
	// 		Action: func(command int) {
	// 			notes.Choose(command, &content.LanguageIndex, len(content.Languages.List))
	// 			content.Language = content.Languages.List[content.LanguageIndex]
	// 		},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.Language.Title
	// 		}},
	// 	&notes.Note{Label: notes.FontSize, Value: notes.IntegerValue,
	// 		Values: []any{layout.FontSize},
	// 		Action: func(command int) {
	// 			notes.Modify(command, &layout.FontSize)
	// 			layout.Refresh(layout.FontSize)
	// 		},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.Layout.FontSize
	// 		}},
	// 	&notes.Note{Label: notes.Monitor, Value: notes.MonitorValue,
	// 		Values: []any{content.Environment.Monitor},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.Environment.Monitor
	// 		},
	// 	},
	// 	&notes.Note{Label: notes.View, Value: notes.ViewValue,
	// 		Values: []any{content.screenWidth, content.screenHeight},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.screenWidth
	// 			values[1] = content.screenHeight
	// 		}},
	// 	&notes.Note{Label: notes.Duration, Value: notes.DurationValue,
	// 		Values: []any{content.CaptureDuration},
	// 		Action: func(command int) {
	// 			notes.ModifyMore(command, &content.CaptureDuration)
	// 		},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.CaptureDuration
	// 		}},
	// 	&notes.Note{Label: notes.FrameRate, Value: notes.IntegerValue, Values: []any{content.currentFrameRate},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.currentFrameRate
	// 		}})
}

func buildCaptureNotes(gs *gizzmo.Game) {
	// content := &gs.Content
	// list := make([]*notes.Note, 0)
	// list = append(list,
	// 	&notes.Note{Label: notes.Capture, Value: notes.CaptureValue,
	// 		Values: []any{content.captureTotal, content.captureCount},
	// 		Refresh: func(values ...any) {
	// 			values[0] = content.captureTotal
	// 			values[1] = content.captureCount
	// 		}})
	// content.CaptureNotes = notes.NewNotes(list,
	// 	content.Languages.List[content.LanguageIndex].Locale)
}

func buildKeyNotes(gs *gizzmo.Game) {
}

func buildPadNotes(gs *gizzmo.Game) {
}
