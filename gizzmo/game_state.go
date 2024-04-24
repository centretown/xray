package gizzmo

import (
	"image/color"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/message"
	msg "github.com/centretown/xray/message"
	"gopkg.in/yaml.v3"
)

const (
	msg_font_size = 20
	msg_label_X   = msg_font_size + 3
	msg_value_X   = msg_label_X + msg_font_size*15
	msg_Y         = msg_font_size + 3
)

var (
	msg_color_label_input = color.RGBA{255, 255, 0, 255}
	msg_color_value_input = color.RGBA{0, 255, 255, 255}
	msg_color_label       = color.RGBA{128, 128, 0, 255}
	msg_color_value       = color.RGBA{0, 128, 128, 255}
)

func (gs *Game) UpdateTokens() {
	content := &gs.Content
	content.monitorNum = rl.GetCurrentMonitor()
	content.monitorWidth = rl.GetMonitorWidth(content.monitorNum)
	content.monitorHeight = rl.GetMonitorHeight(content.monitorNum)
	content.monitorRefreshRate = rl.GetMonitorRefreshRate(content.monitorNum)
	content.currentFrameRate = int64(rl.GetFPS())
	content.screenWidth = int64(rl.GetScreenWidth())
	content.screenHeight = int64(rl.GetScreenHeight())
}

func (gs *Game) BuildTokens() {
	content := &gs.Content
	content.monitorNum = rl.GetCurrentMonitor()
	content.monitorWidth = rl.GetMonitorWidth(content.monitorNum)
	content.monitorHeight = rl.GetMonitorHeight(content.monitorNum)
	content.screenWidth = int64(rl.GetScreenWidth())
	content.screenHeight = int64(rl.GetScreenHeight())
	content.monitorRefreshRate = rl.GetMonitorRefreshRate(content.monitorNum)
	content.currentFrameRate = int64(rl.GetFPS())

	tokens := make([]*msg.Token, 0)
	tokens = append(tokens,
		&msg.Token{Label: msg.Monitor, Value: msg.MonitorValue,
			Values: []any{content.monitorNum,
				content.monitorWidth,
				content.monitorHeight,
				content.monitorRefreshRate},
			Update: func(values ...any) {
				for i := range values {
					switch i {
					case 0:
						values[i] = content.monitorNum
					case 1:
						values[i] = content.monitorWidth
					case 2:
						values[i] = content.monitorHeight
					case 3:
						values[i] = content.monitorRefreshRate
					}
				}
			},
		},
		&msg.Token{Label: msg.View, Value: msg.ViewValue,
			Values: []any{content.screenWidth, content.screenHeight},
			Update: func(values ...any) {
				for i := range values {
					switch i {
					case 0:
						values[i] = content.screenWidth
					case 1:
						values[i] = content.screenHeight
					}
				}
			}},
		&msg.Token{Label: msg.Duration, Value: msg.DurationValue,
			Values: []any{content.CaptureDuration},
			Update: func(values ...any) {
				for i := range values {
					switch i {
					case 0:
						values[i] = content.CaptureDuration
					}
				}
			}},
		&msg.Token{Label: msg.FrameRate, Value: msg.FrameRateValue, Values: []any{content.currentFrameRate},
			Update: func(values ...any) {
				for i := range values {
					switch i {
					case 0:
						values[i] = content.currentFrameRate
					}
				}
			}},
	)
	content.tokens = msg.NewTokenList(tokens)

	tokens = make([]*msg.Token, 0)
	tokens = append(tokens,
		&msg.Token{Label: msg.Capture, Value: msg.CaptureValue,
			Values: []any{content.captureTotal, content.captureCount},
			Update: func(values ...any) {
				for i := range values {
					switch i {
					case 0:
						values[i] = content.captureTotal
					case 1:
						values[i] = content.captureCount
					}
				}

			}})

	content.captureTokens = msg.NewTokenList(tokens)

}

func (gs *Game) DrawStatus() {
	content := &gs.Content
	gs.UpdateTokens()

	if !content.commandState && !content.capturing {
		return
	}

	content.tokens.Fetch()

	row := int(msg_Y)
	row += gs.drawOutputs(row, content.tokens)

	if content.capturing {
		content.captureTokens.Fetch()
		gs.drawOutputs(row, content.captureTokens)
	}
}

func (gs *Game) drawOutputs(row int, tkl *message.TokenList) int {
	var (
		content                  = &gs.Content
		color_label, color_value color.RGBA
	)

	draw := func(i int, label, value string) {
		if i == content.currentToken {
			color_label = msg_color_label_input
			color_value = msg_color_value_input

		} else {
			color_label = msg_color_label
			color_value = msg_color_value
		}
		rl.DrawText(label, int32(msg_label_X), int32(row), msg_font_size, color_label)
		rl.DrawText(value, int32(msg_value_X), int32(row), msg_font_size, color_value)
	}

	for i := range tkl.Outputs {
		tkl.Draw(i, draw)
		row += msg_font_size * 2
	}
	return row
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.SetViewPort(float32(rl.GetRenderWidth()),
		float32(rl.GetRenderHeight()))

	for _, mover := range gs.Content.movers {
		mover.Refresh(current, rl.Vector4{
			X: viewPort.Width,
			Y: viewPort.Height})
	}
	for _, drawer := range gs.Content.drawers {
		drawer.Refresh(current, rl.Vector4{
			X: float32(viewPort.Width),
			Y: float32(viewPort.Height)})
	}
}

func (gs *Game) SetViewPort(rw, rh float32) rl.Rectangle {
	gs.Content.Width = rw
	gs.Content.Height = rh
	return gs.GetViewPort()
}

func (gs *Game) GetViewPort() rl.Rectangle {
	return rl.Rectangle{
		X:      0,
		Y:      0,
		Width:  gs.Content.Width,
		Height: gs.Content.Height,
	}
}

func (game *Game) Dump() {
	buf, _ := yaml.Marshal(game)
	log.Println(string(buf))

	for _, mv := range game.Actors() {
		buf, _ = yaml.Marshal(mv)
		log.Println(string(buf))
		buf, _ = yaml.Marshal(mv.GetDrawer())
		log.Println(string(buf))
	}

	for _, dr := range game.Drawers() {
		buf, _ = yaml.Marshal(dr)
		log.Println(string(buf))
	}
}
