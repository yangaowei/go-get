package extractors

import (
	"../logs"
	"../utils"
	//"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type HuYa struct {
	Base
	vid string
}

// type HuYaItem struct {
// 	Format string `json:"format"`
// }

func HuYaRegister() {
	self := new(HuYa)
	self.Name = "huya"
	self._VIDEO_PATTERNS = []string{`v\.huya\.com\/play\/(\d+)\.html`, `ahuya\.duowan\.com\/play\/(\d+)\.html`}
	Spiders[self.Name] = self
	self.Hd = map[string]string{"350": "normal", "1000": "hd1", "1300": "hd2", "yuanhua": "hd3"}
	//fmt.Println(youku.Name)
}

func (self *HuYa) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()

	self.vid = utils.R1Of(self._VIDEO_PATTERNS, url)
	//html, _ := utils.GetContent(url, nil)
	logs.Log.Debug("vid %s", self.vid)
	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	createTimeStr := utils.R1(`发布于\ (\d+-\d+-\d+ \d+:\d+)`, html)
	logs.Log.Debug("createTimeStr %s", createTimeStr)
	var title string
	var createTime int64
	var duration int64
	title = doc.Find("h1.title").Text()
	if len(createTimeStr) > 0 {
		createTime = utils.StringToMilliseconds("2006-01-02 15:04", createTimeStr)
	}
	infoUrl := fmt.Sprintf("http://v-api-open.huya.com/index.php?r=app/video&vid=%s", self.vid)
	data, _ := self.BuildJson(infoUrl)
	items, _ := data.Get("result").Get("items").Map()
	downloadInfo := make(map[string]interface{})
	for key, value := range items {
		//logs.Log.Debug("key %s,%v", key, value.(map[string]interface{}))
		if hd, ok := self.Hd[key]; ok {
			item := value.(map[string]interface{})
			transcode := item["transcode"].(map[string]interface{})
			urls := transcode["urls"].([]interface{})
			downloadInfo[hd] = map[string]interface{}{"urls": []string{urls[0].(string)}}
			if duration == 0 {
				d, _ := strconv.ParseFloat(item["duration"].(string), 10)
				duration = int64(d)
			}
		}
	}
	info.downloadInfo = downloadInfo
	info.title = title
	info.url = url
	info.duration = duration
	info.createTime = createTime
	return
}

func (self *HuYa) Obj() (obj interface{}) {
	return self
}
