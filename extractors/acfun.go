package extractors

import (
	"../logs"
	"../utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	//simplejson "github.com/bitly/go-simplejson"
)

type AcFun struct {
	Base
	vid string
}

func AcFunRegister() {
	self := new(AcFun)
	self.Name = "acfun"
	self._VIDEO_PATTERNS = []string{`www\.acfun\.(?:tv|cn)\/v\/ac(\d+)`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *AcFun) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			logs.Log.Error("get video info error:%s", err)
		}
	}()
	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	logs.Log.Debug("html %d", len(html))
	vids := utils.FindSubAll(`data-vid=\"(\d+)\"`, html)
	vid := vids[1]
	timeSpan := doc.Find(".time").Text()[:17]
	timeText := strings.Replace(timeSpan, "年", "-", 1)
	timeText = strings.Replace(timeText, "月", "-", 1)
	timeText = strings.Replace(timeText, "日", "", 1)
	timeText = strings.Replace(timeText, " ", "0", 1)
	createTime := utils.StringToMilliseconds("2006-01-02", timeText)
	logs.Log.Debug("timSpan %s", timeText)
	logs.Log.Debug("vid %s", vid)
	var title string
	var duration int64
	title, _ = doc.Find("#pageInfo").Attr("data-title")

	i, _ := self.BuildJson(fmt.Sprintf("http://www.acfun.cn/video/getVideo.aspx?id=%s", vid))
	sourceType, _ := i.Get("sourceType").String()
	logs.Log.Debug("sourceType %s", sourceType)
	info.title = title
	info.url = url
	info.duration = duration / 1000
	info.createTime = createTime
	//info.downloadInfo = tmp
	return
}

func (self *AcFun) Obj() (obj interface{}) {
	return self
}
