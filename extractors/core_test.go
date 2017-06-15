package extractors

import (
	//"../utils"
	"encoding/json"
	"log"
	"testing"
)

func TestYoukuGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["youku"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://v.youku.com/v_show/id_XMjc5NTY4MzMxMg==.html"); err == nil {
		//log.Println(videoInfo)
		info := videoInfo.dumps()
		//log.Println(info)
		//utils.FJson(info)
		b, _ := json.Marshal(info)
		log.Println(string(b))

	} else {
		log.Println(err)
	}
}

func TestQQGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["qq"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://v.qq.com/page/z/7/k/z0389i9cj7k.html"); err == nil {
		//log.Println(videoInfo)
		info := videoInfo.dumps()
		//log.Println(info)
		//utils.FJson(info)
		b, _ := json.Marshal(info)
		log.Println(string(b))

	} else {
		log.Println(err)
	}
}
