package converter

import (
	"log"
	"testing"
)

func TestFFMpeg(t *testing.T) {
	ffmpeg := FFMpeg{}
	log.Println(ffmpeg.which("ffmpeg"))
	options := make(map[string]interface{})
	options["format"] = "mp4"
	audio := make(map[string]string)
	audio["codec"] = "copy"
	options["audio"] = audio

	video := make(map[string]string)
	video["codec"] = "copy"
	video["faststart"] = "true"
	options["video"] = video
	log.Println(ffmpeg.parseOptions(options))

	ffmpeg.Merge([]string{"../download/test_0.mp4", "../download/test_1.mp4"}, "../download/output.mp4", options)
	//log.Println(XOk("/usr/local/bin/ffmpeg"))
}
