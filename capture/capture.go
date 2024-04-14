package capture

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ericpauley/go-quantize/quantize"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func NextFileName(ext, path string) string {
	var namePrefix = "capture"
	id := uuid.New()
	name := fmt.Sprintf("%s_%s.%s", namePrefix, id.String(), ext)
	return filepath.Join(path, name)
}

func createFile(ext string, path string) (io.WriteCloser, error) {
	fname := NextFileName(ext, path)

	w, err := os.Create(fname)
	if err != nil {
		log.Println("Create", fname, err)
	}
	return w, err
}

func CaptureMPG() {
	err := ffmpeg.Input("./sample_data/in1.mp4").
		Output("./sample_data/out1.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return
	}
}

func CapturePNG(path string, img image.Image) {
	w, err := createFile("png", path)
	if err != nil {
		return
	}
	defer w.Close()
	png.Encode(w, img)
}

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

func CaptureGIF(path string, done <-chan int,
	img <-chan image.Image,
	pal color.Palette,
	delay int,
	colorMap map[color.Color]uint8) {

	log.Println("Capturing...")

	var pics = make([]image.Image, 0)
	for {
		select {

		case pic := <-img:
			pics = append(pics, pic)

		case <-done:
			log.Println("Writing...")
			WriteGIF(path, pics, pal, colorMap, delay)
			log.Println("Done.")
			return

		default:
			time.Sleep(time.Millisecond)
		}
	}
}

func PackImages(imgs []image.Image) (combined image.Image) {
	// for _, img := range imgs {
	// 	pt := img.Bounds().Size()
	// 	width, height := pt.X, pt.Y
	// }

	return
}

// https://unix.stackexchange.com/questions/40638/how-to-do-i-convert-an-animated-gif-to-an-mp4-or-mv4-on-the-command-line
// ffmpeg -i animated.gif -movflags faststart -pix_fmt yuv420p \
//		-vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" video.mp4
// movflags – This option optimizes the structure of the MP4 file so the browser can load it as quickly as possible.
// pix_fmt – MP4 videos store pixels in different formats. We include this option to specify a specific format which has maximum compatibility across all browsers.
// vf – MP4 videos using H.264 need to have a dimensions that are divisible by 2. This option ensures that’s the case.
// Source: http://rigor.com/blog/2015/12/optimizing-animated-gifs-with-html5-video

func WriteGIF(path string, pics []image.Image, pal color.Palette,
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

	w, err := createFile("gif", path)
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
		log.Println("EncodeAll", err)
	}
}

func ExtendPalette(colors []color.RGBA, imgs []image.Image,
	count int) (color.Palette, map[color.Color]uint8) {

	return extendPalette(colors, imgs, count)
}

func extendPalette(fixedColors []color.RGBA, imgs []image.Image,
	count int) (color.Palette, map[color.Color]uint8) {
	var (
		bigImg        *image.RGBA
		newPal        = make(color.Palette, 0, count)
		quant         = quantize.MedianCutQuantizer{}
		colorMap      = make(map[color.Color]uint8)
		colorCountMap = make(map[color.RGBA]int)
		// ok         bool
		pixelCount    int
		nonBlankCount int
	)

	for _, c := range fixedColors {
		newPal = append(newPal, c)
	}

	var (
		min, max   int = 5000, 0
		none, rgba color.RGBA
	)

	// rgbaOr := func(condition bool, vals [2]color.RGBA) color.RGBA {
	// 	return vals[check.As[int](!condition)]
	// }

	for _, img := range imgs {

		width, height := img.Bounds().Size().X, img.Bounds().Size().Y
		// t.Logf("width:%d,height:%d", width, height)
		pixelCount += width * height

		for y := range height {
			for x := range width {

				r, g, b, a := img.At(x, y).RGBA()
				rgba = color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				}

				if rgba != none {
					nonBlankCount++
					count = colorCountMap[rgba] + 1
					colorCountMap[rgba] = count
					if count < min {
						min = count
					} else if count > max {
						max = count
					}
				}
			}
		}
	}

	// pixel line
	bigImg = image.NewRGBA(image.Rect(0, 0, 1, nonBlankCount))
	column := 0
	for rgba, count = range colorCountMap {
		for range count {
			bigImg.SetRGBA(0, column, rgba)
		}
		column += count
	}

	// first 256 colors newPal, add to colormap
	newPal = quant.Quantize(newPal, bigImg)
	for v, c := range newPal {
		colorMap[c] = uint8(v)
	}

	// create image with 256 color pal
	paletted := image.NewPaletted(bigImg.Bounds(), newPal)
	model := paletted.ColorModel()
	rect := bigImg.Bounds()

	for y := range rect.Max.Y {
		for x := range rect.Max.X {
			c := bigImg.At(x, y)
			cv := model.Convert(c)
			ix := colorMap[cv]
			colorMap[c] = ix
		}
	}

	// t.Logf("palette length:%d colorMap length=%d",
	// 	len(newPal), len(colorMap))

	return newPal, colorMap
}
