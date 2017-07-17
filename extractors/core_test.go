package extractors

import (
	"../download"
	"../utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		// a := videoInfo.dumps()
		// utils.FJson(a)
		for key, value := range videoInfo.downloadInfo {
			down := (value).(map[string]interface{})
			log.Println(key, down)
			urls := (down["urls"]).([]string)
			log.Println(key, len(urls))
			info := map[string]interface{}{"title": fmt.Sprintf("%s_%s", videoInfo.title, key)}
			result, err := download.DownloadUrls(urls, "mp4", info)
			log.Println(result, err)
		}
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
	if videoInfo, err := ie.GetVideoInfo("http://tv.sohu.com/20170706/n600040516.shtml"); err == nil {
		for key, value := range videoInfo.downloadInfo {
			if key != "normal" {
				continue
			}
			down := (value).(map[string]interface{})
			log.Println(key, down)
			urls := (down["urls"]).([]string)
			log.Println(key, len(urls))
			info := map[string]interface{}{"title": fmt.Sprintf("%s_%s", "sohu", key)}
			result, err := download.DownloadUrls(urls, "mp4", info)

			fmt.Println(result, err)
			break
		}

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

func TestBiLiGetVideoInfo(t *testing.T) {
	//YouKuRegister()
	//url := "http://www.bilibili.com/video/av12112835/"
	url := "http://bangumi.bilibili.com/anime/5832/play#100379"
	key, ie := GetExtractor(url)
	log.Println("get ie:", ie, key)
	if ie != nil {
		if videoInfo, err := ie.GetVideoInfo(url); err == nil {
			for key, value := range videoInfo.downloadInfo {
				if key != "hd2" {
					continue
				}
				down := (value).(map[string]interface{})
				log.Println(key, down)
				urls := (down["urls"]).([]string)
				log.Println(key, len(urls))
				info := map[string]interface{}{"title": fmt.Sprintf("%s_%s", "bilibili", key)}
				header := make(http.Header)
				header.Add("Referer", url)
				info["header"] = header
				result, err := download.DownloadUrls(urls, "mp4", info)

				fmt.Println(result, err)
				break
			}

		} else {
			log.Println(err)
		}
	}
}

func TestBaseGetVideoInfo(t *testing.T) {
	var url string
	//url = "http://bangumi.bilibili.com/anime/5832/play#100379"
	//url = "http://www.toutiao.com/i6428908253530292737/"
	url = "http://v.pptv.com/show/ibZRJxia6UBEKlI4s.html"

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
		utils.FJson(info.Dumps())
	}
}
