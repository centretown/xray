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

func (rayl *RayL) DrawTextEx(font rl.Font, text string, position rl.Vector2,
	fontSize float32, spacing float32, tint color.RGBA) {
	rl.DrawTextEx(font, text, position, fontSize, spacing, tint)
}

func (rayl *RayL) GetRenderWidth() int {
	return rl.GetRenderWidth()
}
func (rayl *RayL) GetRenderHeight() int {
	return rl.GetRenderHeight()
}

func (rayl *RayL) DrawTexture(texture Texture2D, posX int32, posY int32, tint color.RGBA) {
	rl.DrawTexture(rl.Texture2D(texture), posX, posY, tint)
}

func (rayl *RayL) LoadTexture(path string) Texture2D {
	return Texture2D(rl.LoadTexture(path))
}

func (rayl *RayL) UnloadTexture(texture Texture2D) {
	rl.UnloadTexture(rl.Texture2D(texture))
}

func (rayl *RayL) DrawTexturePro(texture Texture2D, sourceRec,
	destRec Rectangle, origin Vector2, rotation float32, tint color.RGBA) {

	rl.DrawTexturePro(rl.Texture2D(texture),
		rl.Rectangle(sourceRec), rl.Rectangle(destRec),
		rl.Vector2(origin), rotation, tint)
}

var rl_ImageClearBackground *rl.Image
var img_ImageClearBackground image.Image

func (rayl *RayL) ImageClearBackground(dst image.Image, c color.RGBA) {
	if dst != img_ImageClearBackground {
		img_ImageClearBackground = dst
		rl_ImageClearBackground = rl.NewImageFromImage(img_ImageClearBackground)
	}
	rl.ImageClearBackground(rl_ImageClearBackground, c)
}
