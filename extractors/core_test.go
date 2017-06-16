package extractors

import (
	"../utils"
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
		utils.FJson(info)
		// b, _ := json.Marshal(info)
		// log.Println(string(b))

	} else {
		log.Println(err)
	}
}

func TestQQGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["qq"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("https://v.qq.com/x/cover/7vofk2li55mtuxe/x0389lut8he.html"); err == nil {
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

func TestIqiyiGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["iqiyi"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://www.iqiyi.com/v_19rr752q8o.html"); err == nil {
		//log.Println(videoInfo)
		info := videoInfo.dumps()
		log.Println(info)
		utils.FJson(info)
		// b, _ := json.Marshal(info)
		// log.Println(string(b))

	} else {
		log.Println(err)
	}
}
