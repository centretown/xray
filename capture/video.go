package capture

import (
	"fmt"
	"image"
	"log"
	"os"
	"syscall"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CaptureVideo(done <-chan int, img <-chan *rl.Image,
	width, height, fps int32) {
	log.Println("CaptureVideo")
	var (
		err    error
		writer *os.File
		// testFile = "./input/capture_e4842eec-044e-4902-becf-30ef508299fd.gif.mp4"
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, writer, err = getPipe()
	if err != nil {
		return
	}
	// return

	// err = ffmpeg.Input(testFile,
	// 	ffmpeg.KwArgs{
	// 		"ss": "1"}).
	// 	Output("./output/"+NextFileName("mkv"),
	// 		ffmpeg.KwArgs{"pix_fmt": "rgba", "t": "20", "r": "25"}).
	// 	OverWriteOutput().ErrorToStdOut().Run()

	// log.Println("CaptureVideo done")

	// time.Sleep(time.Minute)

	resolution := fmt.Sprintf("%dx%d", width, height)

	err = ffmpeg.Input("pipe:0",
		ffmpeg.KwArgs{
			"f":       "rawvideo",
			"pix_fmt": "rgba",
			"s":       resolution,
			"r":       fmt.Sprint(fps),
		}).
		Output("./output/"+NextFileName("mp4"),
			ffmpeg.KwArgs{
				"c:v":     "libx264",
				"vb":      "2500k",
				"pix_fmt": "yuv420p",
			}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		return
	}
	// err = writer.Close()
	// if err != nil {
	// 	return
	// }

	var (
		count      int
		byteCount  int
		frameCount int
		pic        *rl.Image
	)

	write := func(img *rl.Image) (count int, err error) {
		pic, ok := img.ToImage().(*image.RGBA)
		if ok {
			count, err = writer.Write(pic.Pix)
		} else {
			err = fmt.Errorf("wrong image type")
		}
		return
	}

	for {
		select {
		case pic = <-img:
			count, err = write(pic)
			if err != nil {
				return
			}

			byteCount += count
			frameCount++

		case <-done:
			err = writer.Close()
			return

		default:
		}
	}

}

func testVideo() {
	var (
		testFile = "./input/capture_e4842eec-044e-4902-becf-30ef508299fd.gif.mp4"
		// testFile = "./sample_data/in1.mp4"
	)

	// [swscaler @ 0x5641cd0cad80] No accelerated colorspace conversion found from yuv420p to rgb8.

	err := ffmpeg.Input(testFile,
		ffmpeg.KwArgs{
			"ss": "1"}).
		Output("./output/out1.mkv",
			ffmpeg.KwArgs{"pix_fmt": "rgba", "t": "20", "r": "25"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		log.Fatal(err)
	}
}

func getPipe() (reader, writer *os.File, err error) {

	reader, writer, err = os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Dup2(int(reader.Fd()), int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func read(reader *os.File, width, length, stride int) {

	buf := make([]byte, width*length*stride)

	for {
		count, err := reader.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		if count > 0 {
			log.Print(string(buf[:count]))
			// log.Println(count)
		}
		time.Sleep(0)
	}
}

func write(writer *os.File, img <-chan *image.RGBA, stop <-chan int) {
	var (
		count      int
		byteCount  int
		frameCount int
		pic        *image.RGBA
		err        error
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	for {
		select {
		case <-stop:
			err = writer.Close()
			return
		case pic = <-img:
			count, err = writer.Write(pic.Pix)
			if err != nil {
				return
			}

			byteCount += count
			frameCount++

		default:
		}
	}
}
