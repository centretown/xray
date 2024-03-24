package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"time"
	"xray/b2"
	"xray/capture"
	"xray/gpads"
	"xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BLACK uint8 = iota
	RED
	YELLOW
	GREEN
	BLUE
)

var colorsMap = map[color.Color]uint8{
	image.Transparent: BLACK,
	rl.Red:            RED,
	rl.Yellow:         YELLOW,
	rl.Green:          GREEN,
	rl.Blue:           BLUE,
}

var colors = []color.RGBA{
	rl.Black,
	rl.Red,
	rl.Yellow,
	rl.Green,
	rl.Blue,
}

var pal = color.Palette{
	image.Transparent,
	rl.Red,
	rl.Yellow,
	rl.Green,
	rl.Blue,
}

func main() {
	runr := tools.NewRunner(640, 400, 60)
	viewPort := runr.GetViewPort()

	runr.Add(tools.NewBall(30, colors[RED]), tools.NewBouncer(viewPort, 60, 60), 0)
	runr.Add(tools.NewBall(20, colors[YELLOW]), tools.NewBouncer(viewPort, 40, 40), 1)
	runr.Add(tools.NewBall(15, colors[GREEN]), tools.NewBouncer(viewPort, 30, 30), 2)
	runr.Add(tools.NewBall(10, colors[BLUE]), tools.NewBouncer(viewPort, 20, 20), 3)
	Run2d(runr)
}

func Run2d(runr *tools.Runner) {
	runr.SetupWindow("2d")

	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime())
		can_move int32   = 0
		pads             = gpads.NewGPads()
	)

	runr.Refresh(current)

	for !rl.WindowShouldClose() {

		current = rl.GetTime()
		can_move = b2.ToInt32(current > previous+interval)
		previous = float64(can_move) * interval

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for _, run := range runr.Actors {
			run.Animate(can_move, current)
		}
		rl.EndDrawing()

		pads.BeginPad()
		PadInput(pads, current)
	}

	rl.CloseWindow()
}

var (
	next       float64
	polling    bool
	frameCount int
	stopChan   = make(chan int)
	scrChan    = make(chan image.Image)
)

func PadInput(pads *gpads.GPads, current float64) {
	if current > next {
		next = current + .02
		if polling {
			scrChan <- rl.LoadImageFromScreen().ToImage()
			frameCount--
			if frameCount < 0 {
				polling = false
				stopChan <- 1
			}
			return
		}

		for i := range pads.GetStickCount() {
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft) {
				polling = true
				go Poll(stopChan, scrChan)
				frameCount = 360
				return
			}
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight) {
				capture.CapturePNG()
				return
			}
		}
	}
}

func Poll(stop <-chan int, scr <-chan image.Image) {
	var pics = make([]image.Image, 0)
	for {
		select {
		case pic := <-scr:
			pics = append(pics, pic)
		case <-stop:
			fmt.Println("Writing")
			WriteGIF(pics)
			return

		default:
			time.Sleep(0)
			// time.Sleep(time.Millisecond)
		}
	}

}

var fileCounter int

func WriteGIF(pics []image.Image) {
	imageCount := len(pics)
	if imageCount < 1 {
		return
	}

	var images = make([]*image.Paletted, imageCount)
	rect := pics[0].Bounds()

	for i, pic := range pics {
		img := image.NewPaletted(rect, pal)
		for y := range rect.Max.Y {
			for x := range rect.Max.X {
				img.SetColorIndex(x, y, colorsMap[pic.At(x, y)])
			}
		}
		images[i] = img
	}

	fileCounter++
	fname := fmt.Sprintf("/home/dave/src/xray/testimg/cap_gif%d.gif", fileCounter)
	w, err := os.Create(fname)
	if err != nil {
		fmt.Println("Create", fname, err)
		return
	}
	defer w.Close()

	delays := make([]int, imageCount)
	disposals := make([]byte, imageCount)
	for i := range imageCount {
		delays[i] = 4
		disposals[i] = gif.DisposalBackground
	}

	opts := &gif.GIF{
		Image:     images,
		Delay:     delays,
		Disposal:  disposals,
		LoopCount: 0,
		Config: image.Config{
			ColorModel: pal,
			Width:      rect.Dx(),
			Height:     rect.Dy(),
		},
		BackgroundIndex: 0,
	}

	err = gif.EncodeAll(w, opts)
	if err != nil {
		fmt.Println("EncodeAll", fname, err)
	}
}
