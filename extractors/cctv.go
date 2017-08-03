package extractors

import (
	"../logs"
	"../utils"
	//"encoding/json"
	"fmt"
	//"github.com/PuerkitoBio/goquery"
	"strconv"
)

type CCTV struct {
	Base
	vid string
}

func CCTVRegister() {
	self := new(CCTV)
	self.Name = "cctv"
	self._VIDEO_PATTERNS = []string{`tv\.cctv\.com\/\d{4}\/\d{2}\/\d{2}\/(VIDE\w+)\.shtml`}
	Spiders[self.Name] = self
	self.Hd = map[string]string{"lowChapters": "low", "chapters": "normal", "chapters2": "hd1", "chapters3": "hd2", "chapters4": "hd3"}
	//fmt.Println(youku.Name)
}

func (self *CCTV) getVid(url string) (vid string) {
	var html string
	if utils.Match(`tv\.cctv\.com/\d+/\d+/\d+/\w+\.shtml`, url) {
		html, _ = utils.GetContent(url, nil)
		vid = utils.R1(`var guid = "(\w+)"`, html)
	}
	return vid
}

func (self *CCTV) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	self.vid = self.getVid(url)
	logs.Log.Debug("vid %s", self.vid)
	api := fmt.Sprintf("http://vdn.apps.cntv.cn/api/getHttpVideoInfo.do?pid=%s", self.vid)
	data, _ := self.BuildJson(api)
	var title string
	var createTime int64
	var duration int64
	title, _ = data.Get("title").String()
	createTimeStr, _ := data.Get("f_pgmtime").String()
	createTime = utils.StringToMilliseconds("2006-01-02 15:04:04", createTimeStr)
	downloadInfo := make(map[string]interface{})
	video, _ := data.Get("video").Map()
	if duration == 0 {
		d, _ := strconv.ParseFloat(video["totalLength"].(string), 10)
		duration = int64(d)
	}
	for key, value := range video {
		if hd, ok := self.Hd[key]; ok {
			var urls []string
			urlItem := value.([]interface{})
			for _, i := range urlItem {
				logs.Log.Debug("item %v", i)
				item := i.(map[string]interface{})
				urls = append(urls, item["url"].(string))
			}
			downloadInfo[hd] = map[string]interface{}{"urls": urls}
		}
	}
	info.downloadInfo = downloadInfo
	info.title = title
	info.url = url
	info.duration = duration
	info.createTime = createTime
	return
}

func (self *CCTV) Obj() (obj interface{}) {
	return self
}
