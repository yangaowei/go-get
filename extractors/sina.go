package extractors

import (
	"../logs"
	"../utils"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

var duration int64

type Sina struct {
	Base
	vid string
}

func SinaRegister() {
	self := new(Sina)
	self.Name = "sina"
	self._VIDEO_PATTERNS = []string{`video\.sina\.com.cn\/view\/(\d+)\.html`}
	Spiders[self.Name] = self
	//self.Hd = map[string]string{"hc": "normal", "hd": "hd1", "he": "hd2", "sh": "hd3"}
	//fmt.Println(youku.Name)
}

type SinaXml struct {
	Result string `xml:"result"`
	Length int64  `xml:"timelength"`
	Title  string `xml:"vname"`
	Durl   []struct {
		Length int64  `xml:"length"`
		Url    string `xml:"url"`
	} `xml:"durl"`
}

func getK(vid, rand string) (sign string) {
	currentSecond := utils.GetCurrentSeconds()
	s := fmt.Sprintf("%b", currentSecond)
	i, _ := strconv.ParseInt(s[:len(s)-6], 2, 64)
	t := strconv.FormatInt(i, 10)
	sign = utils.MD5(fmt.Sprintf("%sZ6prk18aWxP278cVAH%s%s", vid, t, rand))[:16] + t
	return sign
}

func getDownloadInfoByVid(vid string) (result map[string]interface{}) {
	rand := fmt.Sprintf("0.%d%d", utils.RandInt(10000, 10000000), utils.RandInt(10000, 10000000))
	url := fmt.Sprintf("http://ask.ivideo.sina.com.cn/v_play.php?vid=%s&ran=%s&p=i&k=%s", vid, rand, getK(vid, rand))
	input, _ := utils.GetContent(url, nil)
	//logs.Log.Debug("input %s", input)
	v := SinaXml{}
	err := xml.Unmarshal([]byte(input), &v)
	if err != nil {
		panic(fmt.Sprintf("xml Unmarshal %v", err))
	}
	//logs.Log.Debug("SinaXml %v", v)
	duration = v.Length / 1000
	var urls []string
	for _, item := range v.Durl {
		urls = append(urls, item.Url)
	}
	result = make(map[string]interface{})
	result["normal"] = map[string]interface{}{"urls": urls}
	return
}

func (self *Sina) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	var title string
	var downloadInfo map[string]interface{}
	doc, _ := goquery.NewDocument(url)
	html, _ := doc.Html()
	logs.Log.Debug("html %d", len(html))
	self.vid = utils.R1(`vid:"?(\d+)"?`, html)
	logs.Log.Debug("vid %s", self.vid)
	if len(self.vid) > 0 {
		title = utils.R1(`title\s*:\s*\'([^\']+)\'`, html)
		downloadInfo = getDownloadInfoByVid(self.vid)
	}
	info.url = url
	info.title = title
	info.downloadInfo = downloadInfo
	info.duration = duration
	return
}

func (self *Sina) Obj() (obj interface{}) {
	return self
}
