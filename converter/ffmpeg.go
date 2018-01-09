package converter

import (
	"../utils"
	"fmt"
	//"log"
	"os"
	"path"
	"strings"
)

type FFMpeg struct {
	ffmpegPath string
}

func XOk(filepath string) (result bool) {
	fileInfo, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		//log.Println("this file not exists")
		return
	}
	mode := fileInfo.Mode().String()
	if strings.Index(mode, "x") > 0 {
		result = true
	}
	return
}

func (self *FFMpeg) which(name string) (filepath string) {
	for _, value := range os.Environ() {
		if strings.Index(value, "PATH=") != 0 {
			continue
		}
		paths := strings.Split(value[5:], ":")
		//log.Println(paths)
		for _, p := range paths {
			fp := path.Join(p, name)
			if XOk(fp) {
				filepath = fp
				break
			}
		}
		if len(filepath) > 0 {
			break
		}
	}
	return
}
func (self *FFMpeg) parseOptions(option map[string]interface{}) (options []string) {
	if format, ok := option["format"]; ok {
		options = append(options, []string{"-f", (format).(string)}...)
	}
	if audio, ok := option["audio"]; ok {
		opt_audio := (audio).(map[string]string)
		c := opt_audio["codec"]
		if c == "copy" {
			options = append(options, []string{"-acodec", "copy"}...)
		}
	}
	if video, ok := option["video"]; ok {
		opt_video := (video).(map[string]string)
		c := opt_video["codec"]
		if c == "copy" {
			options = append(options, []string{"-vcodec", "copy"}...)
		}
		if _, ok := opt_video["faststart"]; ok {
			options = append(options, []string{"-movflags", "faststart"}...)
		}
	}
	return
}

func (self *FFMpeg) Merge(videos []string, vfile string, option map[string]interface{}) (result bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ffmpge Merge error: ", err)
		}
		//resp.Body.close()
	}()
	if len(self.ffmpegPath) == 0 {
		self.ffmpegPath = self.which("ffmpeg")
	}
	for _, video := range videos {
		ts := fmt.Sprintf("%s.ts", video)
		cmdParams := fmt.Sprintf("%s -loglevel quiet -y -i '%s' -c copy -f mpegts -bsf:v h264_mp4toannexb %s", self.ffmpegPath, video, ts)
		utils.Cmd(cmdParams)
	}
	concat := self.ffmpegPath + " -loglevel" + " quiet" + "  -y" + " -i" + ` "concat:`
	for _, video := range videos {
		f := fmt.Sprintf("%s.ts", video)
		concat += f + "|"
	}
	concat = concat[:len(concat)-1]
	concat += `"` + " -bsf:a aac_adtstoasc"
	concat += " " + strings.Join(self.parseOptions(option), " ") + " "
	concat += vfile
	//log.Println(concat)
	utils.Cmd(concat)
	result = true
	for _, video := range videos {
		f := fmt.Sprintf("%s.ts", video)
		os.Remove(f)
	}
	return
}

func (self *FFMpeg) Probe(vfile string) (info map[string]interface{}, err error) {
	if _, e := os.Stat(vfile); e != nil {

		return info, e
	}
	return
}
