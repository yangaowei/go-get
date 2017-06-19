package extractors

import (
	"../utils"
	"fmt"
	//"strconv"
	"math/rand"
	"strings"
)

type Sohu struct {
	Base
	Name            string
	_VIDEO_PATTERNS []string
	vid             string
	Hd              map[string]string
}

func SohuRegister() {
	sohu := new(Sohu)
	sohu.Name = "sohu"
	sohu._VIDEO_PATTERNS = []string{}
	Spiders[sohu.Name] = sohu

	sohu.Hd = make(map[string]string)
	//{1: 'normal', 2: 'hd1', 3: 'hd2', 4: 'hd3', 5: 'hd4', 96: 'low'}
	sohu.Hd["norVid"] = "normal"
	sohu.Hd["superVid"] = "hd2"
	sohu.Hd["oriVid"] = "hd2"
	sohu.Hd["relativeId"] = "low"
	sohu.Hd["highVid"] = "hd1"
}

func (self *Sohu) realUrl(url string) string {
	info, _ := self.BuildJson(url)
	s, _ := info.Get("url").String()
	return s

}

func (self *Sohu) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, err := utils.GetContent(url, nil)
	self.vid = utils.R1Of([]string{`\&id=(\d+)`, `vid=\"(\d+)\"`}, html)
	var apiurl string
	if strings.Index(url, "my.tv.sohu.com") > 0 {
		apiurl = "http://my.tv.sohu.com/play/videonew.do?vid=%s&referer=http://my.tv.sohu.com"
	} else {
		apiurl = "http://hot.vrs.sohu.com/vrs_flash.action?vid=%s"
	}
	videoInfo, err := self.BuildJson(fmt.Sprintf(apiurl, self.vid))
	var title string
	var duration int64
	title, _ = videoInfo.Get("data").Get("tvName").String()
	duration = videoInfo.Get("data").Get("totalDuration").MustInt64()

	streamTypes := make(map[string]interface{})
	for key, value := range self.Hd {
		hdvid, _ := videoInfo.Get("data").Get(key).String()
		if hdvid != self.vid {
			videoInfo, err = self.BuildJson(fmt.Sprintf(apiurl, hdvid))
			status := videoInfo.Get("status").MustInt64()
			if status != 1 {
				continue
			}
			host, _ := videoInfo.Get("allot").String()
			//prot, _ := videoInfo.Get("prot").String()
			//tvid, _ := videoInfo.Get("tvid").String()
			su, _ := videoInfo.Get("data").Get("su").Array()
			clipsURL, _ := videoInfo.Get("data").Get("clipsURL").Array()
			ck, _ := videoInfo.Get("data").Get("ck").Array()
			urls := []string{}
			for index, _ := range su {
				n := (su[index]).(string)
				c := (clipsURL[index]).(string)
				c_k := (ck[index]).(string)
				t := rand.Float64()
				url := fmt.Sprintf("http://%s/?prot=9&prod=flash&pt=1&file=%s&new=%s&key=%s&vid=%s&uid=%d&t=%f&rb=1", host, c, n, c_k, self.vid, self.CurrentTime()*1000, t)
				urls = append(urls, self.realUrl(url))
			}
			tmp := make(map[string]interface{})
			tmp["urls"] = urls
			streamTypes[value] = tmp
		}

	}
	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = streamTypes
	return
}
