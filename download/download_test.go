package download

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestUrlSave(t *testing.T) {
	info := make(map[string]interface{})
	info["title"] = "test"
	result, err := DownloadUrls([]string{"http://img.waqu.com/baidutieba/05toxsv7qddqd781.mp4", "http://img.waqu.com/baidutieba/05toxsv7qddqd781.mp4"}, "mp4", info)
	fmt.Println(result)
	fmt.Println(err)
}

func TestDownloadm3u8(t *testing.T) {
	info := make(map[string]interface{})
	info["title"] = "test"
	var urls []string
	m3u8Url := "https://vod-yq.aliyun.com/vod-7651a3/24ea8ccb131642a3820b473f7de594c0/d89fa960d7674b45a3b00923625cb60f-77009f87a3e1473397482d5ec3b2f83a-ld.m3u8"
	m3u8File, err := UrlSave("m3u8.file", m3u8Url, nil)
	if err == nil {
		file, _ := os.Open(m3u8File)
		defer file.Close()
		br := bufio.NewReader(file)
		//dir := path.Dir(m3u8Url)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			line := string(a)
			if !strings.HasPrefix(line, "#") {
				url := "https://vod-yq.aliyun.com/vod-7651a3/24ea8ccb131642a3820b473f7de594c0/" + line
				urls = append(urls, url)
			}
		}
	}
	fmt.Println(Download(urls, "mp4", info))
}
