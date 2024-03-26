package main

import (
	"fmt"
	"image"
	"image/color"
	"xray/b2"
	"xray/capture"
	"xray/gpads"
	"xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WHITE uint8 = iota
	BLACK
	RED
	YELLOW
	GREEN
	CYAN
	BLUE
	MAGENTA
)

var redBall *tools.Ball
var redBouncer *tools.Bouncer

const (
	baseInterval = .02
	screenWidth  = 640
	screenHeight = 480
	fps          = 60
)

func addBalls(runr *tools.Runner, viewPort rl.RectangleInt32, pal color.Palette) {
	redBall = tools.NewBall(40, pal[RED])
	redBouncer = tools.NewBouncer(viewPort, 40, 40)
	runr.Add(tools.NewBall(15, pal[MAGENTA]), tools.NewBouncer(viewPort, 15, 15), 5)
	runr.Add(tools.NewBall(20, pal[BLUE]), tools.NewBouncer(viewPort, 20, 20), 4)
	runr.Add(tools.NewBall(25, pal[CYAN]), tools.NewBouncer(viewPort, 25, 25), 3)
	runr.Add(tools.NewBall(30, pal[GREEN]), tools.NewBouncer(viewPort, 30, 30), 2)
	runr.Add(tools.NewBall(35, pal[YELLOW]), tools.NewBouncer(viewPort, 35, 35), 1)
	runr.Add(redBall, redBouncer, 0)
}

func main() {
	var fixedPal = color.Palette{
		rl.White,
		rl.Black,
		rl.Red,
		rl.Yellow,
		rl.Green,
		color.RGBA{R: 0, G: 255, B: 255, A: 0},
		rl.Blue,
		rl.Magenta,
	}
	runr := tools.NewRunner(screenWidth, screenHeight, fps)
	runr.SetupWindow("2d")
	viewPort := runr.GetViewPort()
	addBalls(runr, viewPort, fixedPal)

	var (
		current  float64 = rl.GetTime()
		previous float64 = current
		interval float64 = float64(rl.GetFrameTime())
		can_move int32   = 0
		pads             = gpads.NewGPads()
	)

	head := rl.LoadTexture("head.png")
	actor := tools.NewHeadText(head)
	bouncer := tools.NewBouncer(viewPort, 60, 60)
	runr.Add(actor, bouncer, 7)

	// pal = capture.ExtendPalette(pal, myface.ToImage())
	// fmt.Println("pal length", len(pal))
	// head := tools.NewHeadText(myface)

	runr.Refresh(current)

	pal, colorMap := createPaletteFromTexture(actor, fixedPal)
	// fmt.Println("Pal", len(pal), pal)

	for !rl.WindowShouldClose() {
		ProcessInput(pads, current, baseInterval, pal, colorMap)

		current = rl.GetTime()
		bMove := current >= previous+interval
		can_move = b2.To[int32](bMove)
		moveFloat := b2.To[float64](bMove)
		previous = moveFloat*interval + moveFloat*current

		if rl.IsWindowResized() {
			runr.Refresh(current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		for _, run := range runr.Actors {
			run.Animate(can_move, current)
		}

		// NOTE: Using DrawTexturePro() we can easily rotate and scale the part of the texture we draw
		// sourceRec defines the part of the texture we use for drawing
		// destRec defines the rectangle where our texture part will fit (scaling it to fit)
		// origin defines the point of the texture used as reference for rotation and scaling
		// rotation defines the texture rotation (using origin as rotation point)
		// DrawTexturePro(scarfy, sourceRec, destRec, origin, (float)rotation, WHITE);

		// x, y := redBouncer.Position()
		// rl.DrawTexturePro(head, x-20, y-20, rl.White)

		rl.EndDrawing()

	}

	rl.UnloadTexture(head)
	rl.CloseWindow()
}

func createPaletteFromTexture(head *tools.HeadText, pal color.Palette) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	head.Draw(0, 0)
	rl.EndDrawing()

	img := rl.LoadImageFromScreen()
	pic := img.ToImage()

	newPal := capture.ExtendPalette(pal, pic)
	colorMap := make(map[color.Color]uint8)
	for v, c := range newPal {
		colorMap[c] = uint8(v)
	}

	p := image.NewPaletted(pic.Bounds(), newPal)
	model := p.ColorModel()
	rect := pic.Bounds()
	for y := range rect.Max.Y {
		for x := range rect.Max.X {
			c := pic.At(x, y)
			cv := model.Convert(c)
			ix := colorMap[cv]
			colorMap[c] = ix
		}
	}

	fmt.Println(colorMap)
	fmt.Println(len(colorMap))
	return newPal, colorMap
}

var (
	next       float64
	capturing  bool
	frameCount int
	stopChan   = make(chan int)
	scrChan    = make(chan image.Image)
)

func ProcessInput(pads *gpads.GPads, current float64, interval float64, pal color.Palette, colorMap map[color.Color]uint8) {
	pads.BeginPad()
	if current > next {
		next = current + interval
		if capturing {
			scrChan <- rl.LoadImageFromScreen().ToImage()
			frameCount--
			if frameCount < 0 {
				capturing = false
				stopChan <- 1
			}
			return
		}

		for i := range pads.GetStickCount() {
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft) {
				capturing = true
				frameCount = 720
				go capture.CaptureGIF(stopChan, scrChan, pal, interval, colorMap)
				// go capture.CaptureGIF(stopChan, scrChan, colorMap, pal)
				return
			}
			if pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight) {
				capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
				return
			}
		}
	}
}
