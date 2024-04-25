package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
)

func (gs *Game) UpdateState(command int) {
	var (
		content = &gs.Content
		length  = content.notes.Length
		note    *notes.Note
	)

	switch command {
	// case notes.TIMES_TEN:
	// 	multiply_by_ten = content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightTrigger1)

	case notes.NEXT_NOTE:
		if content.note+1 < length {
			content.note++
		} else {
			content.note = 0
		}
	case notes.PREV_NOTE:
		if content.note-1 >= 0 {
			content.note--
		} else {
			content.note = length - 1
		}
	case notes.INC, notes.INC_MORE, notes.DEC, notes.DEC_MORE:
		note = content.notes.List[content.note]
		if note.CanAct() {
			note.Action(command)
		}
	case notes.CAPTURE:
		if content.capturing {
			gs.EndCapture()
		} else {
			gs.BeginCapture("mp4")
		}
	case notes.PAUSE_PLAY:
		content.paused = !content.paused
		if !content.paused {
			gs.Refresh(content.currentTime)
		}
	}
}

func (gs *Game) UpdateNotes() {
	content := &gs.Content
	content.monitorNum = rl.GetCurrentMonitor()
	content.monitorWidth = rl.GetMonitorWidth(content.monitorNum)
	content.monitorHeight = rl.GetMonitorHeight(content.monitorNum)
	content.monitorRefreshRate = rl.GetMonitorRefreshRate(content.monitorNum)
	content.currentFrameRate = int64(rl.GetFPS())
	content.screenWidth = int64(rl.GetScreenWidth())
	content.screenHeight = int64(rl.GetScreenHeight())
}

func (gs *Game) BuildNotes() {
	notes.Initialize()
	gs.UpdateNotes()
	content := &gs.Content
	list := make([]*notes.Note, 0)
	list = append(list,
		&notes.Note{Label: notes.Monitor, Value: notes.MonitorValue,
			Values: []any{content.monitorNum,
				content.monitorWidth,
				content.monitorHeight,
				content.monitorRefreshRate},
			UpdateAll: func(values ...any) {
				values[0] = content.monitorNum
				values[1] = content.monitorWidth
				values[2] = content.monitorHeight
				values[3] = content.monitorRefreshRate
			},
		},
		&notes.Note{Label: notes.View, Value: notes.ViewValue,
			Values: []any{content.screenWidth, content.screenHeight},
			UpdateAll: func(values ...any) {
				values[0] = content.screenWidth
				values[1] = content.screenHeight
			}},
		&notes.Note{Label: notes.Duration, Value: notes.DurationValue,
			Values: []any{content.CaptureDuration},
			Action: func(command int) {
				notes.Increment(command, &content.CaptureDuration)
			},
			UpdateAll: func(values ...any) {
				values[0] = content.CaptureDuration
			}},
		&notes.Note{Label: notes.FrameRate, Value: notes.FrameRateValue, Values: []any{content.currentFrameRate},
			UpdateAll: func(values ...any) {
				values[0] = content.currentFrameRate
			}},
	)
	content.notes = notes.NewNotes(list)

	list = make([]*notes.Note, 0)
	list = append(list,
		&notes.Note{Label: notes.Capture, Value: notes.CaptureValue,
			Values: []any{content.captureTotal, content.captureCount},
			UpdateAll: func(values ...any) {
				values[0] = content.captureTotal
				values[1] = content.captureCount
			}})

	content.captureNotes = notes.NewNotes(list)

}
