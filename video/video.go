package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	var (
		err            error
		reader, writer *os.File
		stop           = make(chan int)
		bufChan        = make(chan []byte)
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	// testVideo()
	reader, writer, err = getPipe()
	if err != nil {
		return
	}

	go read(reader, 1920, 1080, 4)
	go write(writer, bufChan, stop)

	now := time.Now()
	til := time.Duration(time.Millisecond * 100)

	for start := time.Duration(0); start < til; start = time.Since(now) {
		bufChan <- []byte(fmt.Sprint(start))
		// time.Sleep(time.Millisecond)
	}

	bufChan <- []byte("world")
	bufChan <- []byte("mamma")

	time.Sleep(time.Millisecond)
	stop <- 1
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
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

func write(writer *os.File, bufchan chan []byte, stop chan int) {
	var (
		count int
		buf   []byte
	)
	for {
		select {
		case <-stop:
			return
		case buf = <-bufchan:
			fmt.Fprintln(writer, string(buf), count)
			count++
		default:
		}
	}
}
