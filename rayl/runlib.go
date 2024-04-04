package rayl

import (
	"image"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectangleInt32 rl.RectangleInt32
type Vector3 rl.Vector3
type Vector2 rl.Vector2
type PixelFormat rl.PixelFormat
type Texture2D rl.Texture2D
type Rectangle rl.Rectangle

var (
	Black   = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	White   = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	Red     = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	Yellow  = color.RGBA{R: 255, G: 255, B: 0, A: 255}
	Green   = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	Cyan    = color.RGBA{R: 0, G: 255, B: 255, A: 255}
	Blue    = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	Magenta = color.RGBA{R: 255, G: 0, B: 255, A: 255}
)

type RunLib interface {
	CloseWindow()
	GetTime() float64
	WindowShouldClose() bool
	IsWindowResized() bool
	BeginDrawing()
	ClearBackground(color.RGBA)
	ImageClearBackground(dst image.Image, color color.RGBA)
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

	LoadTexture(path string) Texture2D
	UnloadTexture(texture Texture2D)
	DrawTexture(texture Texture2D, posX int32, posY int32, tint color.RGBA)
	DrawTexturePro(texture Texture2D, sourceRec, destRec Rectangle, origin Vector2, rotation float32, tint color.RGBA)
}
