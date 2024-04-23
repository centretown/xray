package capture

import (
	"fmt"
	"io"
	"log"

	rl "github.com/centretown/raylib-go/raylib"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CaptureVideo(stop <-chan int, img <-chan *rl.Image,
	width, height, fps int32) {

	log.Println("CaptureVideo")
	var (
		reader, writer = io.Pipe()
		err            error
		fpss           = fmt.Sprintf("%d", fps)
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	go write(stop, img, writer, width, height)
	log.Println("Starting ffmpeg process2")
	done := make(chan error)
	go func() {
		err = ffmpeg.Input("pipe:",
			ffmpeg.KwArgs{
				"format":  "rawvideo",
				"pix_fmt": "rgba",
				"s":       fmt.Sprintf("%dx%d", width, height),
			}).
			Output("./output/"+NextFileName("mp4"),
				ffmpeg.KwArgs{
					"pix_fmt":   "yuv420p",
					"vf":        "scale=trunc(iw/2)*2:trunc(ih/2)*2",
					"framerate": fpss,
				}).
			OverWriteOutput().
			WithInput(reader).
			Run()
		log.Println("ffmpeg process2 done")
		done <- err
		close(done)
	}()

}

func write(done <-chan int, imgCh <-chan *rl.Image, writer io.WriteCloser, width, height int32) {
	log.Println("ffmpeg write")

	const COLOR_WIDTH = 4

	var (
		count      int
		byteCount  int
		frameCount int
		err        error
		// pixels     []byte = make([]byte, width*height*COLOR_WIDTH)
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	writePixels := func(img *rl.Image) error {
		pixels := img.ToBytes()
		// *buf = []byte(ret)
		// *buf = []

		// // Color GetImageColor(Image image, int x, int y);
		// for y := range img.Height {
		// 	for x := range img.Width {
		// 		c := rl.GetImageColor(img, x, y)
		// 		offset := y*width*COLOR_WIDTH + x*COLOR_WIDTH
		// 		buf[offset] = c.R
		// 		buf[offset+1] = c.G
		// 		buf[offset+2] = c.B
		// 		buf[offset+3] = c.A
		// 	}
		// }
		count, err = writer.Write(pixels)
		if err != nil {
			return err
		}
		byteCount += count
		frameCount++
		return nil
	}

	for {
		select {
		case img := <-imgCh:
			err = writePixels(img)
			if err != nil {
				return
			}
			rl.UnloadImage(img)

		case <-done:
			fmt.Println("FFMPEG DONE", frameCount, byteCount)
			err = writer.Close()
			return
		}
	}
}
