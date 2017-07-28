package extractors

import (
	"../logs"
	"../utils"
	"encoding/json"
	"fmt"
	//"github.com/PuerkitoBio/goquery"
)

type YinYueTai struct {
	Base
	vid string
}

type JsonResult struct {
	VideoInfo struct {
		CoreVideoInfo struct {
			VideoId        int64  `json:"videoId"`
			VideoName      string `json:"videoName"`
			VideoUrlModels []struct {
				QualityLevel     string `json:"qualityLevel"`
				QualityLevelName string `json:"qualityLevelName"`
				VideoUrl         string `json:"videoUrl"`
				FileSize         int64  `json:"fileSize"`
			} `json:"videoUrlModels"`
		} `json:"coreVideoInfo"`
	} `json:"videoInfo"`
}

func YinYueTaiRegister() {
	self := new(YinYueTai)
	self.Name = "yinyuetai"
	self._VIDEO_PATTERNS = []string{`(?:v|www)\.yinyuetai\.com\/video\/(\d+)`}
	Spiders[self.Name] = self
	self.Hd = map[string]string{"hc": "normal", "hd": "hd1", "he": "hd2", "sh": "hd3"}
	//fmt.Println(youku.Name)
}

func (self *YinYueTai) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	self.vid = utils.R1Of(self._VIDEO_PATTERNS, url)
	logs.Log.Debug("vid %s", self.vid)
	infoUrl := fmt.Sprintf("http://www.yinyuetai.com/insite/get-video-info?json=true&videoId=%s", self.vid)
	data, _ := utils.GetContent(infoUrl, nil)
	var result JsonResult
	json.Unmarshal([]byte(data), &result)
	//var title string
	var createTime int64
	html, _ := utils.GetContent(url, nil)
	pubTimeStr := utils.R1(`发布于(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})`, html)
	logs.Log.Debug("pubTimeStr %s", pubTimeStr)
	createTime = utils.StringToMilliseconds("2006-01-02 15:04:05", pubTimeStr)
	info.createTime = createTime
	info.title = result.VideoInfo.CoreVideoInfo.VideoName
	info.url = url
	stringsHd := make(map[string]interface{})
	for _, value := range result.VideoInfo.CoreVideoInfo.VideoUrlModels {
		stringsHd[self.Hd[value.QualityLevel]] = map[string]interface{}{"urls": []string{value.VideoUrl}}
	}
	info.downloadInfo = stringsHd

	return
}

func (self *YinYueTai) Obj() (obj interface{}) {
	return self
}
