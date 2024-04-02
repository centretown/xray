package rayl

import (
	"image"
	"image/color"
)

type RunLib interface {
	CloseWindow()
	GetTime() float64
	WindowShouldClose() bool
	IsWindowResized() bool
	BeginDrawing()
	ClearBackground(color.RGBA)
	EndDrawing()
	LogWarnings()
	InitWindow(screenWidth, screenHeight int, title string)
	SetWindowResizble()
	SetTargetFPS(int32)
	GetFPS() int32
	GetFrameTime() float64
	LoadImageFromScreen() image.Image

	GetCurrentMonitor() int
	GetMonitorRefreshRate(monitor int) int
	GetMonitorWidth(monitor int) int
	GetMonitorHeight(monitor int) int
	GetScreenWidth() int
	GetScreenHeight() int
	GetRenderWidth() int
	GetRenderHeight() int

	DrawLine(x, y int32, x1, y1 int32, c color.RGBA)
	DrawCircle(x, y int32, r float32, c color.RGBA)
	DrawText(text string, x, y int32, fontsize int32, c color.RGBA)
}

type RectangleInt32 struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}
