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
	qq._VIDEO_PATTERNS = []string{`v\.qq\.com/page/\w/\w/\w/(\w+)\.html`, `v\.qq\.com/\w/cover/\w{15}/(\w+)\.html`}
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
	vid := utils.R1Of(self._VIDEO_PATTERNS, url)
	log.Println("vid: ", vid)
	api_url := "https://h5vv.video.qq.com/getinfo?callback=&charge=0&vid=" + vid + "&defaultfmt=auto&otype=json&guid=&platform=11001&defnpayver=0&appVer=3.0.16&sdtfrom=v3010&host=m.v.qq.coms&defn=auto&fhdswitch=0&show1080p=0&isHLS=0&fmt=auto&newplatform=11001"
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

	var title string
	var duration int64
	var fn string
	var fvkey string
	stringsHd := make(map[string]interface{})
	if m, ok := (vi).(map[string]interface{}); ok {
		title = (m["ti"]).(string)
		fn = (m["fn"]).(string)
		fvkey = (m["fvkey"]).(string)
		n := (m["td"]).(string)
		duration, _ = strconv.ParseInt(utils.R1(`(\d+)\.`, n), 10, 64)
		ul, _ := json.Marshal(m["ul"])
		downUrl := utils.R1(`(http.+?)\",`, string(ul))
		downUrl = downUrl + fn + "?vkey=" + fvkey
		tmp := make(map[string]interface{})
		urls := []string{downUrl}
		tmp["urls"] = urls
		stringsHd["normal"] = tmp
	}

	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = stringsHd
	return
}

func (self *QQ) Obj() (obj interface{}) {
	return self
}
