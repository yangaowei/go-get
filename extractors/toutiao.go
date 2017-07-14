package extractors

import (
	"../logs"
	"../utils"
	"fmt"
	//"log"
	//"reflect"
	"encoding/base64"
	//"strconv"
	//"strings"
)

type TouTiao struct {
	Base
	vid string
}

func TouTiaoRegister() {
	self := new(TouTiao)
	self.Name = "toutiao"
	self._VIDEO_PATTERNS = []string{`toutiao\.com\/item\/(\d+)`, `toutiao\.com\/[ia](\d+)`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *TouTiao) GetVid(url string) (vid string) {
	if doc, err := self.BuildDoc(url); err == nil {
		fmt.Println(doc)
	}
	return
}

func (self *TouTiao) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	var title string
	var duration int64
	var createTime int64
	title = utils.R1(`title: \'(.+)\'`, html)
	self.vid = utils.R1(`videoid\:\'(\w+)\'`, html)
	createTimeStr := utils.R1(`time: \'(\d+\/\d+\/\d+)\'`, html)
	if len(createTimeStr) > 0 {
		createTime = utils.StringToMilliseconds("2006/01/02", createTimeStr)
	}
	logs.Log.Debug("createTime %s", createTime)
	keyUrl := fmt.Sprintf("http://i.snssdk.com/video/urls/1/toutiao/mp4/%s", self.vid)
	vInfo, _ := self.BuildJson(keyUrl)
	base64Url, _ := vInfo.Get("data").Get("video_list").Get("video_1").Get("main_url").String()
	ts, _ := base64.StdEncoding.DecodeString(base64Url)
	logs.Log.Debug("download url %s", ts)
	stringsHd := map[string]interface{}{"normal": map[string]interface{}{"urls": []string{string(ts)}}}
	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = stringsHd
	info.createTime = createTime

	return
}

func (self *TouTiao) Obj() (obj interface{}) {
	return self
}
