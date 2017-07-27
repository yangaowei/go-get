package extractors

import (
	"../logs"
	"../utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type PearVideo struct {
	Base
	vid string
}

func PearVideoRegister() {
	self := new(PearVideo)
	self.Name = "pearvideo"
	self._VIDEO_PATTERNS = []string{`www\.pearvideo\.com/video_(\d+)`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *PearVideo) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	logs.Log.Debug("html %d", len(html))
	pubTime := doc.Find("div.date").Text()
	var title string
	var createTime int64
	createTime = utils.StringToMilliseconds("2006-01-02 15:04", pubTime)
	title = doc.Find("h1.video-tt").Text()
	info.createTime = createTime
	info.title = title
	info.url = url

	downloadInfo := make(map[string]interface{})

	hdUrl := utils.R1(`hdUrl=\"(.+?)\"`, html)
	sdUrl := utils.R1(`sdUrl=\"(.+?)\"`, html)
	ldUrl := utils.R1(`ldUrl=\"(.+?)\"`, html)
	if len(hdUrl) > 10 {
		downloadInfo["hd1"] = map[string]interface{}{"urls": []string{hdUrl}}
	}
	if len(sdUrl) > 10 {
		downloadInfo["normal"] = map[string]interface{}{"urls": []string{sdUrl}}
	}
	if len(ldUrl) > 10 {
		downloadInfo["low"] = map[string]interface{}{"urls": []string{ldUrl}}
	}
	info.downloadInfo = downloadInfo
	return
}

func (self *PearVideo) Obj() (obj interface{}) {
	return self
}
