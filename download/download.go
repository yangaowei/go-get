package download

import (
	"../converter"
	"../utils"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func UrlSave(vfile, url string) (result string, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err) // 这里的err其实就是panic传入的内容，55
		}
		//resp.Body.close()
	}()
	log.Println("downloading ", vfile)
	for i := 0; i < 3; i++ {
		_, resp := utils.Urlopen(url)
		contentLength, _ := strconv.ParseInt(resp.Header["Content-Length"][0], 10, 64)
		f, _ := os.Create(vfile)
		io.Copy(f, resp.Body)
		if fileInfo, err := os.Stat(vfile); err == nil {
			fileLength := fileInfo.Size()
			if contentLength == fileLength {
				result = vfile
				break
			} else {
				log.Println("file size not equal")
			}
		} else {
			log.Println("vfile not exists ", vfile)
		}
	}
	return
}

func DownloadUrls(urls []string, ext string, info map[string]string) (vfile string, err error) {
	title := info["title"]
	vfile = title + "." + ext
	if len(urls) == 1 {
		vfile, _ = UrlSave(vfile, urls[0])
	} else {
		var vfiles []string
		for index, url := range urls {
			f := fmt.Sprintf("%s_%d.%s", title, index, ext)
			vf, _ := UrlSave(f, url)
			if len(vf) > 0 {
				vfiles = append(vfiles, vf)
			} else {
				panic(fmt.Sprintf("download %s fail", f))
			}
		}
		if len(vfiles) == len(urls) {
			options := map[string]interface{}{"format": "mp4"}
			audio := map[string]string{"codec": "copy"}
			options["audio"] = audio
			video := map[string]string{"codec": "copy", "faststart": "true"}
			options["video"] = video
			conv := converter.FFMpeg{}
			result := conv.Merge(vfiles, vfile, options)
			if !result {
				err = errors.New("Merge videos error")
			}
			for _, vfile := range vfiles {
				os.Remove(vfile)
			}
		}
	}
	return
}
