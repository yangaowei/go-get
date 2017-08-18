package extractors

import (
	"../logs"
	"../utils"
	//"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

type Open163 struct {
	Base
	vid string
}

func Open163Register() {
	self := new(Open163)
	self.Name = "163"
	self._VIDEO_PATTERNS = []string{`open\.163\.com\/movie\/\d{4}\/\d{1,2}\/\w\/\w\/(.+)\.html`}
	Spiders[self.Name] = self
}

func (self *Open163) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()

	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	html = mahonia.NewDecoder("gbk").ConvertString(html)
	var title string
	var createTime int64
	var duration int64
	title = mahonia.NewDecoder("gbk").ConvertString(doc.Find("span.sname").Text())
	downloadInfo := make(map[string]interface{})
	urls := utils.R1Of([]string{`["\'](.+)-list.m3u8["\']`, `["\'](.+).m3u8["\']`}, html) + ".mp4"
	logs.Log.Debug("urls %s", urls)
	downloadInfo["normal"] = map[string]interface{}{"urls": []string{urls}}
	info.downloadInfo = downloadInfo
	info.title = title
	info.url = url
	info.duration = duration
	info.createTime = createTime
	return
}

func (self *Open163) Obj() (obj interface{}) {
	return self
}
