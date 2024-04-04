package rayl

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestFont(t *testing.T) {
	rl.InitWindow(400, 400, "test")
	a := rl.LoadFont("../font/arial.ttf")
	b := rl.LoadFont("../font/DejaVuSans.ttf")
	c := rl.LoadFont("../font/Go-Medium-Italic.ttf")
	d := rl.LoadFont("../font/Roboto-Medium.ttf")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawTextEx(a, "Hello World", rl.Vector2{X: 0, Y: 0}, 40, 4, rl.Red)
		rl.DrawTextEx(b, "Hello World", rl.Vector2{X: 0, Y: 80}, 40, 2, rl.Yellow)
		rl.DrawTextEx(c, "Hello World", rl.Vector2{X: 0, Y: 160}, 40, 2, rl.Green)
		rl.DrawTextEx(d, "Hello World", rl.Vector2{X: 0, Y: 240}, 40, 2, rl.SkyBlue)
		rl.DrawText("Hello World", 0, 320, 40, rl.Orange)

		rl.EndDrawing()
	}

	rl.UnloadFont(a)
	rl.UnloadFont(b)
	rl.UnloadFont(c)
	rl.UnloadFont(d)
	rl.CloseWindow()

	// t.Log(font.Index('t'))
}
