package tools

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	leftMargin   = 20
	rightMargin  = 20
	topMargin    = 50
	bottomMargin = 20
)

type Runner struct {
	width  int32
	height int32
	fps    int32

	actors []*Actor
}

func NewRunner(width int32, height int32, fps int32) *Runner {
	runr := &Runner{
		height: height,
		width:  width,
		fps:    fps,
		actors: make([]*Actor, 0),
	}
	return runr
}

func (runr *Runner) Add(d CanDraw, a CanAnimate, after float64) {
	runr.actors = append(runr.actors, NewActor(d, a, after))
}

func (runr *Runner) Run2d() {
	runr.setupWindow("2d")

	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime())
		can_move int32   = 0
	)

	runr.Refresh(current)

	for !rl.WindowShouldClose() {
		current = rl.GetTime()
		can_move = Bool2int32(current > previous+interval)
		previous = float64(can_move) * interval

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		rl.BeginDrawing()
		runr.setupBackground()

		for _, run := range runr.actors {
			run.Animate(can_move, current)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
	fmt.Println("Were not done yet! Slowly but surely.")
}

func (runr *Runner) Refresh(current float64) {
	viewPort := runr.GetViewPort()
	for _, run := range runr.actors {
		run.Resize(viewPort, current)
	}
}

func (runr *Runner) setupWindow(title string) {
	rl.SetTraceLogLevel(rl.LogInfo)
	rl.InitWindow(runr.width, runr.height, title)
	rl.SetTargetFPS(runr.fps)
	rl.SetWindowState(rl.FlagWindowResizable)
}

func (runr *Runner) setupBackground() {
	runRect := runr.GetViewPort()
	rl.ClearBackground(rl.Black)
	rl.DrawRectangleGradientV(0, 0, int32(rl.GetRenderWidth()), int32(rl.GetRenderHeight()), rl.DarkBlue, rl.Black)
	rl.DrawRectangleGradientV(runRect.X, runRect.Y, runRect.Width, runRect.Height, rl.Black, rl.DarkBlue)
}

func (runr *Runner) Run3d() {
	runr.setupWindow("3d")
	runr.JoyStick(nil)

	camPos := rl.Vector3{X: 10, Y: 10, Z: 10}
	camTar := rl.Vector3{X: 0, Y: 0, Z: 0}
	camUp := rl.Vector3{X: 0, Y: .5, Z: 0}
	camera := rl.NewCamera3D(camPos, camTar, camUp,
		60, rl.CameraPerspective)
	cubeV := rl.Vector3{X: 0, Y: 0, Z: 0}

	for !rl.WindowShouldClose() {
		runr.KeyPosXYZ(&cubeV, &camPos)

		camera.Position = camPos
		rl.BeginDrawing()

		runr.setupBackground() //runRect)
		rl.BeginMode3D(camera)
		rl.DrawGrid(4, 3)
		rl.DrawCubeWires(cubeV, 3, 4, 5, rl.Magenta)
		rl.EndMode3D()

		rl.EndDrawing()
	}
	rl.CloseWindow()
	fmt.Println("THREE D.")
}

func (runr *Runner) JoyStick(obj *rl.Vector3) {
	fmt.Print("game pad ")
	if rl.IsGamepadAvailable(0) {
		fmt.Println("available")
	} else {
		fmt.Println("un-available")
	}
}

func (runr *Runner) KeyPosXYZ(obj, pos *rl.Vector3) {
	x, y, z := rl.IsKeyDown(rl.KeyX), rl.IsKeyDown(rl.KeyY), rl.IsKeyDown(rl.KeyZ)
	up := rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyRight)
	down := rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyLeft)

	vecs := []*rl.Vector3{obj, pos}
	i := Bool2int(rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift))
	v := vecs[i]

	const delta = .25
	v.X -= Bool2float32(up && x) * delta
	v.X += Bool2float32(down && x) * delta
	v.Y += Bool2float32(up && y) * delta
	v.Y -= Bool2float32(down && y) * delta
	v.Z -= Bool2float32(up && z) * delta
	v.Z += Bool2float32(down && z) * delta
}

func (runr *Runner) GetViewPort() rl.RectangleInt32 {
	rw := rl.GetRenderWidth()
	if rw > 0 {
		return rl.RectangleInt32{
			X:      leftMargin,
			Y:      topMargin,
			Width:  int32(rw - leftMargin - rightMargin),
			Height: int32(rl.GetRenderHeight() - topMargin - bottomMargin)}
	}

	return rl.RectangleInt32{
		X:      leftMargin,
		Y:      topMargin,
		Width:  runr.width - leftMargin - rightMargin,
		Height: runr.height - topMargin - bottomMargin,
	}
}
