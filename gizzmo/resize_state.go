package gizzmo

import (
	"fmt"

	rl "github.com/centretown/raylib-go/raylib"
)

type ResizeState int

const (
	RESIZE_SMALL ResizeState = iota
	RESIZE_NORMAL
	RESIZE_FULLSCREEN
	RESIZE_MINIMUM = RESIZE_SMALL
)

func (gs *Game) resize(state ResizeState) {
	content := &gs.Content
	if state < RESIZE_MINIMUM {
		state = RESIZE_MINIMUM
	} else if state > RESIZE_FULLSCREEN {
		state = RESIZE_FULLSCREEN
	}
	if state == content.screenstate {
		return
	}

	fmt.Println("state", state)
	if state == RESIZE_FULLSCREEN {
		content.Width = content.FixedWidth
		content.Height = content.FixedHeight
	} else {
		if content.fullscreen {
			content.Width = content.FixedWidth
			content.Height = content.FixedHeight
		}

		switch content.screenstate {
		case RESIZE_MINIMUM:
			content.Width = content.FixedWidth / 2
			content.Height = content.FixedHeight / 2
		default:
			content.Width = content.FixedWidth
			content.Height = content.FixedHeight
		}
		rl.SetWindowSize(int(content.Width), int(content.Height))
	}

	rl.UnloadRenderTexture(gs.Content.renderTexture)
	content.Width = float32(rl.GetRenderWidth())
	content.Height = float32(rl.GetRenderHeight())
	content.renderTexture = rl.LoadRenderTexture(
		int32(rl.GetRenderWidth()),
		int32(rl.GetRenderHeight()))
	gs.Refresh(content.currentTime)
}
