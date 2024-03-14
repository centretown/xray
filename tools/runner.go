package tools

import (
	"fmt"
	"image/color"

	"xray/b2i"
	"xray/gpads"

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
	gpads  *gpads.GPads
}

func NewRunner(width int32, height int32, fps int32) *Runner {
	runr := &Runner{
		height: height,
		width:  width,
		fps:    fps,
		actors: make([]*Actor, 0),
	}

	runr.gpads = gpads.NewGPads()

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
		can_move = b2i.Bool2int32(current > previous+interval)
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
}

func (runr *Runner) AddBouncingBalls() {
	var colors = []color.RGBA{
		rl.White,
		rl.Blue,
		rl.Yellow,
		rl.Red,
		rl.White,
		rl.Red,
		rl.White,
		rl.Yellow,
		rl.Lime,
		rl.DarkGreen,
	}

	viewPort := runr.GetViewPort()
	viewPort.Height -= viewPort.Height / 2

	runr.Add(NewBall(60, colors), NewBouncer(viewPort, 60, 60), 0)
	runr.Add(NewBall(40, colors[6:]), NewBouncer(viewPort, 40, 40), 1)
	runr.Add(NewBall(30, colors[2:]), NewBouncer(viewPort, 30, 30), 2)
	runr.Add(NewBall(20, colors[4:]), NewBouncer(viewPort, 20, 20), 3)

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
	// runRect := runr.GetViewPort()
	rl.ClearBackground(rl.Black)
	// rl.DrawRectangleGradientV(0, 0, int32(rl.GetRenderWidth()), int32(rl.GetRenderHeight()), rl.DarkBlue, rl.Black)
	// rl.DrawRectangleGradientV(runRect.X, runRect.Y, runRect.Width, runRect.Height, rl.Black, rl.DarkBlue)
}

type draw3dFunc struct {
	draw func(pos rl.Vector3, color color.RGBA)
}

var drawCubeWires = draw3dFunc{
	draw: func(pos rl.Vector3, color color.RGBA) {
		rl.DrawCubeWires(pos, 4, 4, 4, color)
	},
}

var drawCube = draw3dFunc{
	draw: func(pos rl.Vector3, color color.RGBA) {
		var (
			position = rl.Vector3{X: pos.X, Y: pos.Y, Z: pos.Z + 5}
		)
		rl.DrawCubeV(position, rl.Vector3{X: 4, Y: 4, Z: 4}, color)
	},
}

var drawGrid = draw3dFunc{
	draw: func(rl.Vector3, color.RGBA) {
		rl.DrawGrid(6, 2)
	},
}

var drawCircle3d = draw3dFunc{
	draw: func(pos rl.Vector3, color color.RGBA) {
		var (
			position         = rl.Vector3{X: pos.X, Y: pos.Y, Z: pos.Z - 5}
			rotation         = rl.Vector3{X: 0, Y: 0, Z: 0}
			angle    float32 = 45.0
		)
		rl.DrawCircle3D(position, 4, rotation, angle, color)
	},
}

func (runr *Runner) Run3d() {
	runr.setupWindow("3d")
	runr.gpads.BeginPad()
	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime() * 100)
		can_move int32   = 0

		camPos    = rl.Vector3{X: 10, Y: 10, Z: 10}
		camTarget = rl.Vector3{X: 0, Y: 0, Z: 0}
		camUp     = rl.Vector3{X: 0, Y: .5, Z: 0}
		camera    = rl.NewCamera3D(camPos, camTarget, camUp,
			60, rl.CameraPerspective)
		cubeV = rl.Vector3{X: 0, Y: 0, Z: 0}

		drawObjects = []draw3dFunc{drawCubeWires, drawCircle3d, drawCube, drawGrid}
	)

	runr.AddBouncingBalls()

	// shader := rl.LoadShader("../shaders/lightint.vs", "../shaders/lightint.fs")

	for !rl.WindowShouldClose() {
		current = rl.GetTime()
		can_move = b2i.Bool2int32(current > previous+interval)
		previous = float64(can_move) * interval

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		runr.gpads.BeginPad()
		runr.PadPosXYZ(&cubeV, &camPos)
		runr.KeyPosXYZ(&cubeV, &camPos)
		camera.Position = camPos

		rl.BeginDrawing()
		{
			runr.setupBackground() //runRect)

			rl.BeginMode3D(camera)
			{
				for _, obj := range drawObjects {
					obj.draw(cubeV, rl.Green)
				}
			}
			rl.EndMode3D()

			rl.DrawCircle(100, 100, 25, rl.Red)
			for _, run := range runr.actors {
				run.Animate(can_move, current)
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
	fmt.Println("THREE D.")
}

func (runr *Runner) PadPosXYZ(obj, pos *rl.Vector3) {
	p := runr.gpads
	count := p.GetStickCount()
	for pi := range count {
		x, y, z := p.GetPadAxisMovement(pi, gpads.ABS_X),
			p.GetPadAxisMovement(pi, gpads.ABS_Y),
			p.GetPadAxisMovement(pi, gpads.ABS_Z)

		const delta float32 = 1.0 / 16.0
		pos.X += delta * x
		pos.Y += delta * y
		pos.Z += delta * z

		px, py := p.GetPadAxisValue(pi, gpads.ABS_HAT0X),
			p.GetPadAxisValue(pi, gpads.ABS_HAT0Y)
		pos.X += float32(px)
		pos.Y += float32(py)
	}
}

func (runr *Runner) KeyPosXYZ(obj, pos *rl.Vector3) {
	x, y, z := rl.IsKeyDown(rl.KeyX), rl.IsKeyDown(rl.KeyY), rl.IsKeyDown(rl.KeyZ)
	up := rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyRight)
	down := rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyLeft)

	vecs := []*rl.Vector3{obj, pos}
	i := b2i.Bool2int(rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift))
	v := vecs[i]

	const delta = .25
	v.X -= b2i.Bool2float32(up && x) * delta
	v.X += b2i.Bool2float32(down && x) * delta
	v.Y += b2i.Bool2float32(up && y) * delta
	v.Y -= b2i.Bool2float32(down && y) * delta
	v.Z -= b2i.Bool2float32(up && z) * delta
	v.Z += b2i.Bool2float32(down && z) * delta
}

func (runr *Runner) GetViewPort() rl.RectangleInt32 {
	rw := rl.GetRenderWidth()
	if rw > 0 {
		return rl.RectangleInt32{
			X:      0,
			Y:      0,
			Width:  int32(rw),
			Height: int32(rl.GetRenderHeight()),
		}
	}

	return rl.RectangleInt32{
		X:      leftMargin,
		Y:      topMargin,
		Width:  runr.width,
		Height: runr.height,
	}
}
