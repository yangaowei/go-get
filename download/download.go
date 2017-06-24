package download

import (
	"../utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func UrlSave(vfile, url string) (result string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err) // 这里的err其实就是panic传入的内容，55
		}
		//resp.Body.close()
	}()
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
		UrlSave(vfile, urls[0])
	} else {
		var vfiles []string
		for index, url := range urls {
			f := fmt.Sprintf("%s_%d.%s", title, index, ext)
			vf := UrlSave(f, url)
			vfiles = append(vfiles, vf)
		}
	}
	return
}
