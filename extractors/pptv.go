package extractors

import (
	"../logs"
	"../utils"
	"fmt"
)

type PPTV struct {
	Base
	vid string
}

func PPTVRegister() {
	self := new(PPTV)
	self.Name = "pptv"
	self._VIDEO_PATTERNS = []string{`v\.pptv\.com\/show\/(\w+)\.html`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *PPTV) GetVid(url string) (vid string) {
	if doc, err := self.BuildDoc(url); err == nil {
		fmt.Println(doc)
	}
	return
}

func (self *PPTV) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	logs.Log.Debug("len %d", len(html))
	vid := utils.R1(`webcfg\s*=\s*{"id":\s*(\d+)`, html)
	logs.Log.Debug("vid %s", vid)
	if len(vid) == 0 {
		panic("vid is null")
	}
	var title string
	var duration int64
	var createTime int64

	info.title = title
	info.url = url
	info.duration = duration
	//info.downloadInfo = stringsHd
	info.createTime = createTime

	return
}

func (self *PPTV) Obj() (obj interface{}) {
	return self
}
