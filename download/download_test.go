package download

import (
	"log"
	//"os"
	"testing"
)

func TestUrlSave(t *testing.T) {
	info := make(map[string]string)
	info["title"] = "test"
	result, err := DownloadUrls([]string{"http://img.waqu.com/baidutieba/05toxsv7qddqd781.mp4", "http://img.waqu.com/baidutieba/05toxsv7qddqd781.mp4"}, "mp4", info)
	log.Println(result)
	log.Println(err)
}
