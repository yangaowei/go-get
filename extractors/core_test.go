package extractors

import (
	"../utils"
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestYoukuGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["youku"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://v.youku.com/v_show/id_XMjgyODc0NTU2MA==.html"); err == nil {
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

func TestIqSohuGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["sohu"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://my.tv.sohu.com/pl/9365360/90147443.shtml"); err == nil {
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

func TestIqLeTvGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	ie := Spiders["letv"]
	log.Println(ie, "Ie")
	if videoInfo, err := ie.GetVideoInfo("http://www.le.com/ptv/vplay/30105085.html"); err == nil {
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

func TestBaseGetVideoInfo(t *testing.T) {
	var url string
	url = os.Args[1]

	log.Println("url:", url)
	log.Println(os.Args)
	var key string
	var spider Core
	for a, b := range Spiders {
		if b.MatchUrl(url) {
			key = a
			spider = b
			break
		}
	}
	log.Println("get IE ", key)
	if len(key) == 0 {
		log.Println("暂不支持该站点")
	} else {
		info, _ := spider.GetVideoInfo(url)
		log.Println(info.Dumps())
	}
}
