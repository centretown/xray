package tools

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ballRadius   = 60
	leftMargin   = 200
	rightMargin  = 20
	topMargin    = 50
	bottomMargin = 20
)

type Runner struct {
	height int32
	width  int32
	fps    int32

	actors []*Runnable
}

func NewRunner() *Runner {
	runr := &Runner{
		height: 720,
		width:  1280,
		fps:    60,
		actors: make([]*Runnable, 0),
	}
	return runr
}

func (runr *Runner) Add(d Drawable, a Animator, start float64, after float64) {
	runr.actors = append(runr.actors, NewRunnable(d, a, start, after))
}

func (runr *Runner) Run(control <-chan int) {

	rl.SetTraceLogLevel(rl.LogInfo)
	rl.InitWindow(runr.width, runr.height, "Runner")
	rl.SetTargetFPS(runr.fps)
	rl.SetWindowState(rl.FlagWindowResizable)

	previous := rl.GetTime()
	current := previous
	interval := float64(rl.GetFrameTime())
	can_move := int32(0)

	var (
		hasElapsed bool
		closeNow   bool
		runRect    = rl.RectangleInt32{
			X:      leftMargin,
			Y:      topMargin,
			Width:  int32(rl.GetRenderWidth() - leftMargin - rightMargin),
			Height: int32(rl.GetRenderHeight() - topMargin - bottomMargin)}
	)

	runr.Add(NewBall(60, colors), NewBouncer(runRect, ballRadius, ballRadius), current, 0)
	runr.Add(NewBall(40, colors[2:]), NewBouncer(runRect, ballRadius, ballRadius), current, 1)
	runr.Add(NewBall(20, colors[3:]), NewBouncer(runRect, ballRadius, ballRadius), current, 2)

	for !closeNow && !rl.WindowShouldClose() {
		current = rl.GetTime()
		can_move = B2int32(current > previous+float64(interval))
		hasElapsed = current > previous+interval
		previous = float64(B2int(hasElapsed)) * interval

		if rl.IsWindowResized() {
			runRect = rl.RectangleInt32{
				X:      leftMargin,
				Y:      topMargin,
				Width:  int32(rl.GetRenderWidth() - leftMargin),
				Height: int32(rl.GetRenderHeight() - topMargin)}

			for _, run := range runr.actors {
				run.Resize(runRect, ballRadius, ballRadius, current)
			}
		}

		rl.BeginDrawing()

		rl.DrawRectangleGradientV(0, 0, int32(rl.GetRenderWidth()), int32(rl.GetRenderHeight()), rl.Maroon, rl.Black)
		rl.DrawRectangleGradientV(runRect.X, runRect.Y, runRect.Width, runRect.Height, rl.SkyBlue, rl.DarkBlue)
		for _, run := range runr.actors {
			run.Animate(can_move, current)
		}
		rl.EndDrawing()

		select {
		case <-control:
			closeNow = true
		default:
		}
	}

	rl.CloseWindow()
	fmt.Println("Were not done yet! Slowly but surely.")

}
