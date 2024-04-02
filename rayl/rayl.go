package rayl

import (
	"image"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RayL struct{}

var _ RunLib = (*RayL)(nil)

func (rayl *RayL) GetTime() float64 {
	return rl.GetTime()
}
func (rayl *RayL) GetFrameTime() float64 {
	return float64(rl.GetFrameTime())
}

func (rayl *RayL) CloseWindow() {
	rl.CloseWindow()
}
func (rayl *RayL) WindowShouldClose() bool {
	return rl.WindowShouldClose()
}
func (rayl *RayL) IsWindowResized() bool {
	return rl.IsWindowResized()
}
func (rayl *RayL) BeginDrawing() {
	rl.BeginDrawing()
}
func (rayl *RayL) ClearBackground(c color.RGBA) {
	rl.ClearBackground(c)
}
func (rayl *RayL) SetTargetFPS(c int32) {
	rl.SetTargetFPS(int32(c))
}

func (rayl *RayL) GetFPS() int32 {
	return rl.GetFPS()
}

func (rayl *RayL) EndDrawing() {
	rl.EndDrawing()
}

func (rayl *RayL) SetWindowResizble() {
	rl.SetWindowState(rl.FlagWindowResizable)
}
func (rayl *RayL) LogWarnings() {
	rl.SetTraceLogLevel(rl.LogWarning)
}

func (rayl *RayL) InitWindow(screenWidth, screenHeight int, title string) {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), title)
}

func (rayl *RayL) LoadImageFromScreen() image.Image {
	return rl.LoadImageFromScreen().ToImage()
}

func (rayl *RayL) GetCurrentMonitor() int {
	return rl.GetCurrentMonitor()
}

func (rayl *RayL) GetMonitorRefreshRate(monitor int) int {
	return rl.GetMonitorRefreshRate(monitor)
}

func (rayl *RayL) GetMonitorWidth(monitor int) int {
	return rl.GetMonitorWidth(monitor)
}
func (rayl *RayL) GetMonitorHeight(monitor int) int {
	return rl.GetMonitorHeight(monitor)
}
func (rayl *RayL) GetScreenWidth() int {
	return rl.GetScreenWidth()
}
func (rayl *RayL) GetScreenHeight() int {
	return rl.GetScreenHeight()
}

func (rayl *RayL) DrawCircle(x, y int32, r float32, c color.RGBA) {
	rl.DrawCircle(x, y, r, c)
}

func (rayl *RayL) DrawLine(x, y int32, x1, y1 int32, c color.RGBA) {
	rl.DrawLine(x, y, x1, y1, c)
}

func (rayl *RayL) DrawText(text string, x int32, y int32, font_size int32, c color.RGBA) {
	rl.DrawText(text, x, y, font_size, c)
}

func (rayl *RayL) GetRenderWidth() int {
	return rl.GetRenderWidth()
}
func (rayl *RayL) GetRenderHeight() int {
	return rl.GetRenderHeight()
}
