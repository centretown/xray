package capture

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"time"

	"github.com/ericpauley/go-quantize/quantize"
)

var fileCounter = 0

func createFile(ext string) (*os.File, error) {
	fileCounter++
	fname := fmt.Sprintf("/home/dave/src/xray/testimg/cap%d.%s", fileCounter, ext)
	w, err := os.Create(fname)
	if err != nil {
		fmt.Println("Create", fname, err)
	}
	return w, err
}

func CapturePNG(img image.Image) {
	w, err := createFile("png")
	if err != nil {
		return
	}
	defer w.Close()
	png.Encode(w, img)
}

// func CaptureGIF(stop <-chan int, scr <-chan image.Image,
// 	colorMap map[color.Color]uint8, pal color.Palette) {
// 	fmt.Println("Capturing...")

func CaptureGIF(stop <-chan int, scr <-chan image.Image, pal color.Palette, interval float64) {
	fmt.Println("Capturing...")

	var pics = make([]image.Image, 0)
	for {
		select {
		case pic := <-scr:
			pics = append(pics, pic)
		case <-stop:
			fmt.Println("Writing...")
			WriteGIF(pics, pal, interval)
			fmt.Println("Done.")
			return

		default:
			time.Sleep(0)
			// time.Sleep(time.Millisecond)
		}
	}

}

func MakePalette(img image.Image) (pal color.Palette, colorMap map[color.Color]uint8) {
	pal = make(color.Palette, 0, 63)
	q := quantize.MedianCutQuantizer{}
	pal = q.Quantize(pal, img)
	colorMap = make(map[color.Color]uint8)
	for i, c := range pal {
		colorMap[c] = uint8(i)
	}
	return
}

func WriteGIF(pics []image.Image, pal color.Palette, interval float64) {
	imageCount := len(pics)
	if imageCount < 1 {
		return
	}

	// interval ms
	// gif delays 10ms

	var delayT int = int(interval * 100)
	fmt.Println("delayT", delayT)

	var images = make([]*image.Paletted, imageCount)
	pic := pics[0]
	rect := pic.Bounds()

	// pal, colorMap := MakePalette(pic)
	colorMap := make(map[color.Color]uint8)
	for v, c := range pal {
		colorMap[c] = uint8(v)
	}

	for i, pic := range pics {
		img := image.NewPaletted(rect, pal)
		for y := range rect.Max.Y {
			for x := range rect.Max.X {
				img.SetColorIndex(x, y, colorMap[pic.At(x, y)])
			}
		}
		images[i] = img
	}

	w, err := createFile("gif")
	if err != nil {
		return
	}
	defer w.Close()

	delays := make([]int, imageCount)
	disposals := make([]byte, imageCount)
	for i := range imageCount {
		delays[i] = delayT
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
		fmt.Println("EncodeAll", w.Name(), err)
	}
}
