package capture

import (
	"fmt"

	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func NextFileName(ext string) string {
	var namePrefix = "capture"
	id := uuid.New()
	name := fmt.Sprintf("%s_%s.%s", namePrefix, id.String(), ext)
	return name
}

func CaptureMPG() {
	err := ffmpeg.Input("./sample_data/in1.mp4").
		Output("./sample_data/out1.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return
	}
}

// https://unix.stackexchange.com/questions/40638/how-to-do-i-convert-an-animated-gif-to-an-mp4-or-mv4-on-the-command-line
// ffmpeg -i animated.gif -movflags faststart -pix_fmt yuv420p \
//		-vf "scale=trunc(iw/2)*2:trunc(ih/2)*2" video.mp4
// movflags – This option optimizes the structure of the MP4 file so the browser can load it as quickly as possible.
// pix_fmt – MP4 videos store pixels in different formats. We include this option to specify a specific format which has maximum compatibility across all browsers.
// vf – MP4 videos using H.264 need to have a dimensions that are divisible by 2. This option ensures that’s the case.
// Source: http://rigor.com/blog/2015/12/optimizing-animated-gifs-with-html5-video
