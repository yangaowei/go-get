package extractors

import (
	"../logs"
	"../utils"
	"crypto/md5"
	"encoding/xml"
	"fmt"
	//"net/http"
	//"reflect"
	//"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	//"strconv"
	//"strings"
)

var (
	appkey               = "f3bb208b3d081dc8"
	SECRETKEY_MINILOADER = "1c15888dc316e05a15fdd0a02ed6584f"
)

type BiLiBiLi struct {
	Base
	vid string
}

func BiLiBiLiRegister() {
	self := new(BiLiBiLi)
	self.Name = "bilibili"
	self._VIDEO_PATTERNS = []string{`www\.bilibili.com\/video\/av(\d+)\/?`, `www\.bilibili\.com\/video\/av(\d+)\/index_\d+\.html`, `bangumi\.bilibili\.com\/anime\/\d+\/play#(\d+)`}
	Spiders[self.Name] = self
	self.Hd = make(map[string]string)
	//'1080p': u'1080p' , '1300': u'超清', '1000': u'高清' , '720p': u'标清', '350': u'流畅'
	self.Hd["3"] = "hd2"
	self.Hd["2"] = "hd1"
	self.Hd["1"] = "normal"
	//fmt.Println(youku.Name)
}

type Video struct {
	Durls         []Durl `xml:"durl"`
	Result        string `xml:"result"`
	Timelength    int64  `xml:"timelength"`
	Format        string `xml:"format"`
	AcceptFormat  string `xml:"accept_format"`
	AcceptQuality string `xml:"accept_quality"`
	From          string `xml:"from"`
	SeekParam     string `xml:"seek_param"`
	SeekType      string `xml:"seek_type"`
}

type Durl struct {
	Order  int64  `xml:"order"`
	Length string `xml:"length"`
	Url    string `xml:"url"`
	Size   int64  `xml:"size"`
}

func parsePlayUrl(input string) (urls []string, ext string, size int64, duration int64, err error) {
	v := Video{}
	err = xml.Unmarshal([]byte(input), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	for _, durl := range v.Durls {
		urls = append(urls, durl.Url)
		size += durl.Size
	}
	ext = v.Format
	duration = v.Timelength
	return
}

func (self *BiLiBiLi) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	timeStr := utils.R1(`\<i\>(\d{4}\-\d{2}\-\d{2}\s\d{2}\:\d{2})\<\/i\>`, html)
	var createTime int64
	if len(timeStr) > 0 {
		createTime = utils.StringToMilliseconds("2006-01-02 15:04", timeStr)
	}
	var title string
	var duration int64
	vid := utils.R1Of([]string{`cid=(\d+)`, `cid=\"(\d+)`}, html)
	//logs.Log.Informational("vid %s", html)
	title = utils.R1Of([]string{`<meta name="title" content="([^<>]{1,999})" />`, `<h1[^>]*>([^<>]+)</h1>`}, html)
	titles := utils.R1(`<option value=.*selected>(.+)</option>`, html)
	if len(vid) == 0 {
		eid := utils.R1Of([]string{`anime/v/(\d+)`, `play#(\d+)`}, url)
		logs.Log.Informational("eids %s", eid)
		eidApi := fmt.Sprintf("http://bangumi.bilibili.com/web_api/episode/%s.json", eid)
		eidHtml, _ := utils.GetContent(eidApi, map[string]interface{}{"proxy": "http://123.59.188.21:8118"})
		bjson := []byte(eidHtml)
		eidInfo, _ := simplejson.NewJson(bjson)
		t, _ := eidInfo.Get("result").Get("currentEpisode").Get("longTitle").String()
		vid, _ = eidInfo.Get("result").Get("currentEpisode").Get("danmaku").String()
		titles = t
	}
	tmp := make(map[string]interface{})
	for key, _ := range self.Hd {
		signStr := fmt.Sprintf("cid=%s&from=miniplay&player=1&quality=%s%s", vid, key, SECRETKEY_MINILOADER)
		h := md5.New()
		h.Write([]byte(signStr))
		sign := fmt.Sprintf("%x", h.Sum(nil))
		apiUrl := fmt.Sprintf("http://interface.bilibili.com/playurl?cid=%s&player=1&quality=%s&from=miniplay&sign=%s", vid, key, sign)
		apiHtml, _ := utils.GetContent(apiUrl, map[string]interface{}{"proxy": "http://123.59.188.21:8118"})
		urls, _, _, d, _ := parsePlayUrl(apiHtml)
		tmp[self.Hd[key]] = map[string]interface{}{"urls": urls}
		duration = d
	}

	info.url = url
	if len(titles) > 0 {
		info.title = title + "_" + titles
	} else {
		info.title = title
	}
	info.duration = duration / 1000
	info.createTime = createTime
	info.downloadInfo = tmp
	return
}

func (self *BiLiBiLi) Obj() (obj interface{}) {
	return self
}
