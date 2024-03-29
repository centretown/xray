package capture

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"os"
	"time"

	"github.com/ericpauley/go-quantize/quantize"
)

var fileCounter = 0
var path = "/home/dave/src/xray/testimg"
var namePrefix = "capture"

func NextFileName(ext string) string {
	return fmt.Sprintf("%s/%s%d.%s", path, namePrefix, fileCounter, ext)
}

func createFile(ext string) (*os.File, error) {
	fname := NextFileName(ext)
	fileCounter++

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

type Cheap struct {
	pal      color.Palette
	colorMap map[color.Color]uint8
}

func NewCheap(pal color.Palette, colorMap map[color.Color]uint8) *Cheap {
	ch := &Cheap{
		pal:      pal,
		colorMap: colorMap,
	}
	return ch
}

func (ch *Cheap) Quantize(p color.Palette, m image.Image) color.Palette {
	return ch.pal
}

func (ch *Cheap) Draw(dst draw.Image, rect image.Rectangle, src image.Image, sp image.Point) {
	for y := range rect.Max.Y {
		for x := range rect.Max.X {
			index := ch.colorMap[src.At(x, y)]
			dst.Set(x, y, ch.pal[index])
		}
	}
}

func WriteGIFFrame(w io.Writer, pic image.Image, cheap *Cheap) {
	gif.Encode(w, pic, &gif.Options{
		NumColors: 64,
		Quantizer: cheap,
		Drawer:    cheap,
	})
}

func CaptureGIF(done <-chan int,
	img <-chan image.Image,
	pal color.Palette,
	delay int,
	colorMap map[color.Color]uint8) {

	fmt.Println("Capturing...")

	var pics = make([]image.Image, 0)
	for {
		select {

		case pic := <-img:
			pics = append(pics, pic)

		case <-done:
			fmt.Println("Writing...")
			WriteGIF(pics, pal, colorMap, delay)
			fmt.Println("Done.")
			return

		default:
			time.Sleep(time.Millisecond)
		}
	}

}

func ExtendPalette(pal color.Palette, img image.Image,
	count int) (color.Palette, map[color.Color]uint8) {

	newPal := make(color.Palette, 0, count)
	newPal = append(newPal, pal...)
	q := quantize.MedianCutQuantizer{}
	newPal = q.Quantize(newPal, img)
	colorMap := make(map[color.Color]uint8)
	for v, c := range newPal {
		colorMap[c] = uint8(v)
	}

	paletted := image.NewPaletted(img.Bounds(), newPal)
	model := paletted.ColorModel()
	rect := img.Bounds()

	for y := range rect.Max.Y {
		for x := range rect.Max.X {
			c := img.At(x, y)
			cv := model.Convert(c)
			ix := colorMap[cv]
			colorMap[c] = ix
		}
	}

	return newPal, colorMap
}

func WriteGIF(pics []image.Image, pal color.Palette,
	colorMap map[color.Color]uint8, delay int) {

	imageCount := len(pics)
	if imageCount < 1 {
		return
	}

	var images = make([]*image.Paletted, imageCount)
	pic := pics[0]
	rect := pic.Bounds()

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
		delays[i] = delay
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
