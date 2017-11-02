package extractors

import (
	"../utils"
	"fmt"
	"log"
	//"reflect"
	"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	"strconv"
	//"strings"
)

type QQ struct {
	Base
	vid string
}

func QQRegister() {
	qq := new(QQ)
	qq.Name = "qq"
	qq._VIDEO_PATTERNS = []string{`v\.qq\.com/page/\w/\w/\w/(\w+)\.html`, `v\.qq\.com/\w/cover/(\w+)\.html`, `v\.qq\.com/\w/cover/\w{15}/(\w+)\.html`}
	Spiders[qq.Name] = qq
	//fmt.Println(youku.Name)
}

func (self *QQ) GetVid(url string) (vid string) {
	if doc, err := self.BuildDoc(url); err == nil {
		fmt.Println(doc)
	}
	return
}

func (self *QQ) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	//vid := utils.R1Of([]string{`"vid":"(\w+)",`}, html)
	vid := "v0568e9mz0s"
	log.Println("vid: ", vid)
	api_url := "http://vv.video.qq.com/getinfo?otype=json&appver=3.2.19.333&platform=11&defnpayver=1&vid=" + vid
	html, gerr := utils.GetContent(api_url, nil)
	if gerr != nil {
		return info, gerr
	}
	video_html := utils.R1("QZOutputJson=(.*)?;", html)
	bjson := []byte(video_html)
	videoInfo, err := simplejson.NewJson(bjson)
	if err != nil {
		return
	}
	vis, _ := videoInfo.Get("vl").Get("vi").Array()
	vi := vis[0]
	vh, _ := (vi).(map[string]interface{})["vh"].(json.Number).Int64()
	var title string
	var duration int64
	// var fn string
	// var fvkey string
	title = (vi).(map[string]interface{})["ti"].(string)
	fnPre := (vi).(map[string]interface{})["lnk"].(string)
	host := (vi).(map[string]interface{})["ul"].(map[string]interface{})["ui"].([]interface{})[0].(map[string]interface{})["url"].(string)
	streams, _ := videoInfo.Get("fl").Get("fi").Array()
	seg_cnt, _ := (vi).(map[string]interface{})["cl"].(map[string]interface{})["fc"].(json.Number).Int64()
	if seg_cnt == 0 {
		seg_cnt = 1
	}
	stringsHd := make(map[string]interface{})
	for _, item := range streams {
		fi := item.(map[string]interface{})
		qualityName := fi["name"].(string)
		qualityId, _ := fi["id"].(json.Number).Int64()
		//fmt.Println(item)
		tmp := make(map[string]interface{})
		var urls []string
		for i := 1; i <= int(seg_cnt); i++ {
			var filename string
			if seg_cnt == 1 && vh <= 480 {
				filename = fnPre + ".mp4"
			} else {
				filename = fmt.Sprintf("%s.p%d.%d.mp4", fnPre, qualityId%10000, i)
			}
			keyApi := fmt.Sprintf("http://vv.video.qq.com/getkey?otype=json&platform=11&format=%d&vid=%s&filename=%s&appver=3.2.19.333", qualityId, vid, filename)
			partInfo, _ := utils.GetContent(keyApi, nil)
			partInfo = utils.R1("QZOutputJson=(.*)?;", partInfo)
			bjson := []byte(partInfo)
			portJson, _ := simplejson.NewJson(bjson)
			vkey, e := portJson.Get("key").String()
			if e == nil {
				downUrl := fmt.Sprintf("%s%s?vkey=%s", host, filename, vkey)
				urls = append(urls, downUrl)
			}
		}
		if len(urls) > 0 {
			tmp["urls"] = urls
			stringsHd[qualityName] = tmp
		}
	}

	n := (vi).(map[string]interface{})["td"].(string)
	duration, _ = strconv.ParseInt(utils.R1(`(\d+)\.`, n), 10, 64)
	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = stringsHd
	return
}

func (self *QQ) Obj() (obj interface{}) {
	return self
}
