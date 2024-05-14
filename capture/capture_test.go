package capture

import (
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/centretown/xray/numbers"
	"github.com/ericpauley/go-quantize/quantize"
)

func TestSort(t *testing.T) {
	var (
		pts = []image.Point{
			{X: 20, Y: 300},
			{X: 203, Y: 200},
			{X: 210, Y: 10},
			{X: 60, Y: 320},
			{X: 15, Y: 350},
			{X: 23, Y: 400},
			{X: 377, Y: 300},
		}
	)

	show(t, pts)

	sorted := sortPtsByHeight(pts)

	showIndex(t, pts, sorted)

}

func sortPtsByHeight(pts []image.Point) []int {

	c := func(a int, b int) int {
		v := numbers.AsOr(pts[a].Y < pts[b].Y, 1, 0)
		return numbers.AsOr(pts[a].Y > pts[b].Y, -1, v)
	}

	srt := make([]int, len(pts))
	for i := range pts {
		srt[i] = i
	}

	slices.SortFunc(srt, c)
	return srt
}

func sortByHeight(imgs []image.Image) (sorted []image.Image) {
	cmp := func(a image.Image, b image.Image) int {
		aY, bY := a.Bounds().Size().Y, a.Bounds().Size().Y
		v := numbers.AsOr(aY < bY, 1, 0)
		return numbers.AsOr(aY > bY, -1, v)
	}

	slices.SortFunc(imgs, cmp)
	return imgs
}

func show[T any](t *testing.T, list []T) {
	for _, i := range list {
		t.Log(i)
	}

}
func showIndex[T any](t *testing.T, list []T, indeces []int) {
	for _, i := range indeces {
		t.Log(list[i])
	}

}

var picsFolder = "test_images"

var pics = []string{
	"doorstop.png",
	"gander.png",
	"head_300.png",
	"moon-solo-300.png",
	"polar.png",
}

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

var fixedPalette = color.Palette{
	White,
	Black,
	Red,
	Yellow,
	Green,
	Cyan,
	Blue,
	Magenta,
}

func TestExtend(t *testing.T) {
	imgs := testLoadImages(t)
	testExtendPalette(t, fixedPalette, imgs, 256)
}

func testExtendPalette(t *testing.T, fixedPal color.Palette, imgs []image.Image,
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

	newPal = append(newPal, fixedPal...)

	var (
		min, max               int = 5000, 0
		none, rgba, cmax, cmin color.RGBA
	)

	// rgbaOr := func(condition bool, vals [2]color.RGBA) color.RGBA {
	// 	return vals[check.As[int](!condition)]
	// }

	for _, img := range imgs {

		width, height := img.Bounds().Size().X, img.Bounds().Size().Y
		t.Logf("width:%d,height:%d", width, height)
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
						cmin = rgba
					} else if count > max {
						max = count
						cmax = rgba
					}
				}
			}
		}
	}

	t.Logf("total pixels/non-blank/colors: %d/%d/%d max=%d (%v) min=%d (%v)",
		pixelCount, nonBlankCount, len(colorCountMap),
		max, cmax, min, cmin)

	// return fixedPalette, colorMap
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

	t.Logf("palette length:%d colorMap length=%d",
		len(newPal), len(colorMap))

	return newPal, colorMap
}

func testLoadImage(t *testing.T, path string) (img image.Image, err error) {
	var (
		title = "LoadImage"
		rdr   io.ReadCloser
		enc   string
	)

	rdr, err = os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	img, enc, err = image.Decode(rdr)
	if err != nil {
		log.Printf("%s %s file %s\n", title, path, err)
	} else {
		log.Printf("%s %s file %s success\n", title, enc, path)
	}
	if err = rdr.Close(); err != nil {
		t.Fatal(err)

	}
	return
}

func testLoadImages(t *testing.T) (imgs []image.Image) {
	imgs = make([]image.Image, len(pics))
	var err error
	for i, pic := range pics {
		imgs[i], err = testLoadImage(t, filepath.Join(picsFolder, pic))
		if err != nil {
			t.Fatal(err)
		}
	}
	return imgs
}
