package main

import (
	"image/color"
	"log"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/check"

	"github.com/centretown/gpads/gpads"

	_ "image/gif"
	_ "image/png"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	runr := NewRunner(1280, 720)
	gpads := gpads.NewGPads()
	Run3d(runr, gpads)
}

type Draw3d struct {
	draw func(pos rl.Vector3, color color.RGBA)
}

var drawCubeWires = Draw3d{
	draw: func(pos rl.Vector3, color color.RGBA) {
		rl.DrawCubeWires(pos, 4, 4, 4, color)
	},
}

var drawCube = Draw3d{
	draw: func(pos rl.Vector3, color color.RGBA) {
		var (
			position = rl.Vector3{X: pos.X, Y: pos.Y, Z: pos.Z + 5}
		)
		rl.DrawCubeV(position, rl.Vector3{X: 4, Y: 4, Z: 4}, color)
	},
}

var drawGrid = Draw3d{
	draw: func(rl.Vector3, color.RGBA) {
		rl.DrawGrid(6, 2)
	},
}

var drawCircle3d = Draw3d{
	draw: func(pos rl.Vector3, color color.RGBA) {
		var (
			position         = rl.Vector3{X: pos.X, Y: pos.Y, Z: pos.Z - 5}
			rotation         = rl.Vector3{X: 0, Y: 0, Z: 0}
			angle    float32 = 45.0
		)
		rl.DrawCircle3D(position, 4, rotation, angle, color)
	},
}

func Run3d(runr *Run3D, gpads *gpads.GPads) {
	rl.InitWindow(1280, 720, "3d")

	gpads.BeginPad()
	gpads.DumpPad()
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

		drawObjects = []Draw3d{drawCubeWires, drawCircle3d, drawCube, drawGrid}
	)

	runr.Refresh(current)

	// shader := rl.LoadShader("../shaders/lightint.vs", "../shaders/lightint.fs")

	for !rl.WindowShouldClose() {
		current = rl.GetTime()
		can_move = check.As[int32](current > previous+interval)
		previous = float64(can_move) * interval

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		gpads.BeginPad()
		PadPosXYZ(gpads, &cubeV, &camPos, current)
		// KeyPosXYZ(&cubeV, &camPos)

		camera.Position = camPos

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode3D(camera)
		for _, obj := range drawObjects {
			obj.draw(cubeV, rl.Green)
		}
		rl.EndMode3D()

		rl.DrawCircle(100, 100, 25, rl.Red)
		for _, run := range runr.Actors {
			run.Move(true, current)
		}
		rl.EndDrawing()

	}
	rl.CloseWindow()
	log.Println("THREE D.")
}

var nextTime float64

func PadPosXYZ(gpad *gpads.GPads, obj, pos *rl.Vector3, current float64) {
	count := gpad.GetPadCount()
	for pi := range count {
		x, y := gpad.GetGamepadAxisMovement(pi, rl.GamepadAxisLeftX),
			gpad.GetGamepadAxisMovement(pi, rl.GamepadAxisLeftY)

		const delta float32 = 1.0 / 16.0
		pos.X += delta * x
		pos.Y += delta * y

		x, y = gpad.GetGamepadAxisMovement(pi, rl.GamepadAxisRightX),
			gpad.GetGamepadAxisMovement(pi, rl.GamepadAxisRightY)
		obj.X += delta * x
		obj.Y -= delta * y

		obj.Z += check.As[float32]((gpad.IsGamepadButtonDown(pi, rl.GamepadButtonLeftFaceDown))) / 4
		obj.Z -= check.As[float32]((gpad.IsGamepadButtonDown(pi, rl.GamepadButtonLeftFaceUp))) / 4
		obj.X += check.As[float32]((gpad.IsGamepadButtonDown(pi, rl.GamepadButtonLeftFaceRight))) / 4
		obj.X -= check.As[float32]((gpad.IsGamepadButtonDown(pi, rl.GamepadButtonLeftFaceLeft))) / 4
		home := check.As[float32](gpad.IsGamepadButtonUp(pi, rl.GamepadButtonRightFaceLeft))
		obj.X, obj.Y, obj.Z = home*obj.X, home*obj.Y, home*obj.Z

		if current > nextTime && gpad.IsGamepadButtonDown(pi, rl.GamepadButtonMiddleLeft) {
			capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
			nextTime = current + .5
		}
	}
}

func KeyPosXYZ(obj, pos *rl.Vector3) {
	x, y, z := rl.IsKeyDown(rl.KeyX), rl.IsKeyDown(rl.KeyY), rl.IsKeyDown(rl.KeyZ)
	up := rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyRight)
	down := rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyLeft)

	vecs := []*rl.Vector3{obj, pos}
	i := check.As[int](rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift))
	v := vecs[i]

	const delta = .25
	v.X -= check.As[float32](up && x) * delta
	v.X += check.As[float32](down && x) * delta
	v.Y += check.As[float32](up && y) * delta
	v.Y -= check.As[float32](down && y) * delta
	v.Z -= check.As[float32](up && z) * delta
	v.Z += check.As[float32](down && z) * delta
}
