package extractors

import (
	"../utils"
	"crypto/sha1"
	"fmt"
	"log"
	//"reflect"
	//"encoding/json"
	//simplejson "github.com/bitly/go-simplejson"
	//"strconv"
	"strings"
)

type LeTv struct {
	Base
	vid string
}

func LeTvRegister() {
	self := new(LeTv)
	self.Name = "letv"
	self._VIDEO_PATTERNS = []string{`www\.(?:le|letv)\.com\/ptv\/vplay\/(\d+)\.html`}
	Spiders[self.Name] = self
	self.Hd = make(map[string]string)
	//'1080p': u'1080p' , '1300': u'超清', '1000': u'高清' , '720p': u'标清', '350': u'流畅'
	self.Hd["1080p"] = "hd4"
	self.Hd["1300"] = "hd3"
	self.Hd["1000"] = "hd2"
	self.Hd["720p"] = "hd1"
	self.Hd["350"] = "normal"
	//fmt.Println(youku.Name)
}

func calcTimeKey(ts int64) (timeKey int64) {
	magic := 185025305
	val := ts
	a := val & (2<<31 - 1)
	l := a >> 8
	r := (val << (32 - 8) & (2<<31 - 1))
	timeKey = (l | r) ^ int64(magic)
	return
}

func decode(data string) (m3u8 string) {
	version := data[0:5]
	if strings.ToLower(version) == "vc_01" {
		loc2 := data[5:]
		length := len(loc2)
		loc4 := make([]int, length*2)
		log.Println(length, len(loc4))
		for i := 0; i < length; i++ {
			loc4[2*i] = int(loc2[i]) >> 4
			loc4[2*i+1] = int(loc2[i]) & 15
		}
		loc6 := append(loc4[len(loc4)-11:], loc4[:len(loc4)-11]...)
		loc7 := make([]int, length)
		for i := 0; i < length; i++ {
			loc7[i] = (loc6[2*i] << 4) + loc6[2*i+1]
		}
		for i := 0; i < len(loc7); i++ {
			m3u8 += fmt.Sprintf("%c", rune(loc7[i]))
		}

	} else {
		m3u8 = data
	}
	return
}

func (self *LeTv) GetVid(url string) (vid string) {
	vid = utils.R1Of(self._VIDEO_PATTERNS, url)
	return
}

func (self *LeTv) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	self.vid = self.GetVid(url)
	apiUrl := fmt.Sprintf("http://player-pc.le.com/mms/out/video/playJson?id=%s&platid=1&splatid=101&format=1&tkey=%d&domain=www.le.com&region=cn&source=1000&accesyx=1", self.vid, calcTimeKey(self.CurrentTime()))
	videoInfo, err := self.BuildJson(apiUrl)
	hostUrls, _ := videoInfo.Get("msgs").Get("playurl").Get("domain").Array()
	hostUrl := (hostUrls[0]).(string)
	log.Println("hostUlr:", hostUrl)

	streams, _ := videoInfo.Get("msgs").Get("playurl").Get("dispatch").Map()
	streamTypes := make(map[string]interface{})
	for key, value := range streams {
		streamUrls := (value).([]interface{})
		sUrl := hostUrl + (streamUrls[0]).(string)
		h := sha1.New()
		h.Write([]byte(sUrl))
		uuid := fmt.Sprintf("%x_0", h.Sum(nil))
		sUrl = strings.Replace(sUrl, "tss=0", "tss=ios", -1)
		sUrl += fmt.Sprintf("&m3v=1&termid=1&format=1&hwtype=un&ostype=MacOS10.12.4&p1=1&p2=10&p3=-&expect=3&tn=&vid=%s&uuid=%s&sign=letv", self.vid, uuid)
		//log.Println("streamUrls:", sUrl)
		streamInfo, _ := self.BuildJson(sUrl)
		suffix := fmt.Sprintf("&r=%d&appid=500", self.CurrentTime()*1000)
		m3u8, _ := streamInfo.Get("location").String()
		m3u8 += suffix
		m3u8Html, _ := utils.GetContent(m3u8, nil)
		m3u8List := decode(m3u8Html)
		urls := utils.FindAll(`http\:\/\/.+`, m3u8List)
		tmp := make(map[string][]string)
		tmp["urls"] = urls
		streamTypes[self.Hd[key]] = tmp
	}
	//log.Println("hostUlr:", streams)

	var title string
	var duration int64
	title, _ = videoInfo.Get("msgs").Get("playurl").Get("title").String()
	duration = videoInfo.Get("msgs").Get("playurl").Get("duration").MustInt64()
	// fmt.Println("duration", duration)
	// fmt.Println("err", err)
	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = streamTypes
	return
}

func (self *LeTv) Obj() (obj interface{}) {
	return self
}
