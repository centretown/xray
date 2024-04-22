package capture

import (
	"fmt"
	"image"
	"io"
	"log"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CaptureVideo(stop <-chan int, img <-chan *image.RGBA,
	width, height, fps int32) {

	log.Println("CaptureVideo")
	var (
		reader, writer = io.Pipe()
		err            error
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	go write(stop, img, writer)
	log.Println("Starting ffmpeg process2")
	done := make(chan error)
	go func() {
		err = ffmpeg.Input("pipe:",
			ffmpeg.KwArgs{"format": "rawvideo",
				"pix_fmt": "rgba", "s": fmt.Sprintf("%dx%d", width, height),
			}).
			Output("./output/"+NextFileName("mp4"), ffmpeg.KwArgs{"pix_fmt": "yuv420p"}).
			OverWriteOutput().
			WithInput(reader).
			Run()
		log.Println("ffmpeg process2 done")
		done <- err
		close(done)
	}()

}

func write(done <-chan int, img <-chan *image.RGBA, writer io.WriteCloser) {
	log.Println("ffmpeg write")

	var (
		count      int
		byteCount  int
		frameCount int
		err        error
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case pic := <-img:
			count, err = writer.Write(pic.Pix)
			byteCount += count
			frameCount++

		case <-done:
			fmt.Println("FFMPEG DONE", frameCount, byteCount)
			err = writer.Close()
			return
		}
	}
}
