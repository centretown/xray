package capture

import (
	"fmt"
	"image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"

	_ "image/gif"
)

var fileCounter = 1

func CapturePNG() {
	img := rl.LoadImageFromScreen().ToImage()
	fname := fmt.Sprintf("/home/dave/src/xray/testimg/cap%d.png", fileCounter)
	w, err := os.Create(fname)
	if err != nil {
		fmt.Println("Create", fname, err)
		return
	}

	defer w.Close()
	png.Encode(w, img)
	fileCounter++
}
