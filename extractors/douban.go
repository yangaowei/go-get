package extractors

import (
	"../logs"
	"../utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type DouBan struct {
	Base
	vid string
}

func DouBanRegister() {
	self := new(DouBan)
	self.Name = "douban"
	self._VIDEO_PATTERNS = []string{`movie\.douban\.com\/trailer\/(\d+)(?:|/#content)`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *DouBan) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	var title string
	logs.Log.Debug("html %d", len(html))
	self.vid = utils.R1Of(self._VIDEO_PATTERNS, url)
	logs.Log.Debug("vid %s", self.vid)
	title = doc.Find("#content>h1").Text()
	timeStr := doc.Find(".trailer-info>span").Text()[:10]
	logs.Log.Debug("timeStr %s", timeStr)
	downloadInfo := make(map[string]interface{})

	downloadInfo["normal"] = map[string]interface{}{"urls": []string{fmt.Sprintf("https://movie.douban.com/trailer/video_url?tid=%s&hd=0", self.vid)}}
	var createTime int64
	createTime = utils.StringToMilliseconds("2006-01-02", timeStr)
	info.title = title
	info.url = url
	// info.duration = duration / 1000
	info.createTime = createTime
	info.downloadInfo = downloadInfo
	return
}

func (self *DouBan) Obj() (obj interface{}) {
	return self
}
