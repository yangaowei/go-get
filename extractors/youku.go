package extractors

import (
	"../utils"
	"fmt"
	//"log"
	//"reflect"
	"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	"strconv"
	"strings"
)

type YouKu struct {
	Base
	vid string
}

func YouKuRegister() {
	youku := new(YouKu)
	youku.Name = "youku"
	youku._VIDEO_PATTERNS = []string{`youku\.com/v_show/id_([a-zA-Z0-9=]+)`}
	Spiders[youku.Name] = youku
	youku.Hd = make(map[string]string)
	//{'3gphd': 'normal', 'hd3': 'hd3', 'hd2': 'hd2', 'mp4hd3': 'hd3','mp4hd2': 'hd2', 'flv': 'normal', 'mp4hd': 'hd1', 'mp4': 'hd1', 'flvhd': 'normal'}
	youku.Hd["3gphd"] = "normal"
	youku.Hd["hd3"] = "hd3"
	youku.Hd["hd2"] = "hd2"
	youku.Hd["mp4hd3"] = "hd3"
	youku.Hd["mp4hd2"] = "hd2"
	youku.Hd["flv"] = "normal"
	youku.Hd["flvhd"] = "normal"
	youku.Hd["mp4hd"] = "hd1"

	//fmt.Println(youku.Name)
}

func fetch_cna() (cna string) {
	url := "http://gm.mmstat.com/yt/ykcomment.play.commentInit?cna="
	_, resp := utils.Urlopen(url)
	cookies := resp.Header["Set-Cookie"]
	cna = utils.R1("cna=([^;]+)", strings.Join(cookies, ";"))
	if len(cna) == 0 {
		cna = "oqikEO1b7CECAbfBdNNf1PM1"
	}
	// //return cna if cna else "oqikEO1b7CECAbfBdNNf1PM1"
	return
}

func (self *YouKu) GetVid(url string) (vid string) {
	if doc, err := self.BuildDoc(url); err == nil {
		fmt.Println(doc)
	}
	return
}

func (self *YouKu) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	vid := utils.R1Of(self._VIDEO_PATTERNS, url)
	//log.Println("vid:", vid)
	self.vid = vid
	//cna := fetch_cna()
	api_url := fmt.Sprintf("https://ups.youku.com/ups/get.json?vid=%s&ccode=%s&client_ip=192.168.1.1&utid=%s&client_ts=%d", self.vid, "0401", fetch_cna(), self.CurrentTime())
	data := make(map[string]interface{})
	html, gerr := utils.GetContent(api_url, data)
	if gerr != nil {
		return info, gerr
	}
	bjson := []byte(html)
	videoInfo, err := simplejson.NewJson(bjson)
	if err != nil {
		return
	}
	streams, _ := videoInfo.Get("data").Get("stream").Array()
	var duration int64
	//var title string
	stringsHd := make(map[string]interface{})
	for _, stream := range streams {
		tmp := make(map[string]interface{})
		if m, ok := (stream).(map[string]interface{}); ok {
			hd := (m["stream_type"]).(string)
			if v, ok := self.Hd[hd]; ok {
				hd = v
			}
			tmp["m3u8_url"] = m["m3u8_url"]
			//duration = (m["milliseconds_video"]).(int64)
			n := (m["milliseconds_video"]).(json.Number)
			duration, _ = strconv.ParseInt(string(n), 10, 64)
			segs, _ := (m["segs"]).([]interface{})
			urls := []string{}
			for _, seg := range segs {
				s, _ := (seg).(map[string]interface{})
				urls = append(urls, (s["cdn_url"]).(string))
				//fmt.Println(s["cdn_url"])
			}
			tmp["urls"] = urls
			stringsHd[hd] = tmp
		}
		//break
	}
	//utils.FJson(stringsHd)
	// log.Println(string(sjson))
	title, _ := videoInfo.Get("data").Get("video").Get("title").String()
	info.title = title
	info.url = url
	info.duration = duration / 1000
	info.downloadInfo = stringsHd
	return
}

func (self *YouKu) Obj() (obj interface{}) {
	return self
}
