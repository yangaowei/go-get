package extractors

import (
	"../logs"
	"../utils"
	"encoding/json"
	"fmt"
	//"net/url"
	// "github.com/PuerkitoBio/goquery"
	// "strconv"
)

type ZuiYou struct {
	Base
	vid string
}

type ZuiYouJson struct {
	Data struct {
		Post struct {
			Status  int    `json:"status"`
			Content string `json:"content"`
			Videos  struct {
				Video struct {
					Url      string `json:"url"`
					UrlSrc   string `json:"urlsrc"`
					UrlExt   string `json:"urlext"`
					Duration int64  `json:"dur"`
				} `json:"102383040"`
			} `json:"videos"`
		} `json:"post"`
	} `json:"data"`
}

func ZuiYouRegister() {
	self := new(ZuiYou)
	self.Name = "izuiyou"
	self._VIDEO_PATTERNS = []string{`www\.izuiyou\.com\/detail\/(\d+)`}
	Spiders[self.Name] = self
	//fmt.Println(youku.Name)
}

func (self *ZuiYou) GetVideoInfo(u string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	self.vid = utils.R1Of(self._VIDEO_PATTERNS, u)
	logs.Log.Debug("vid %s", self.vid)
	value := fmt.Sprintf(`{"pid":%s}`, self.vid)
	html, err := utils.PostContent("http://www.izuiyou.com/api/post/detail", nil, value)
	var s ZuiYouJson
	json.Unmarshal([]byte(html), &s)
	fmt.Println("s:", s)

	info.title = s.Data.Post.Content
	info.duration = s.Data.Post.Videos.Video.Duration
	info.url = u
	tmp := make(map[string]interface{})
	tmp["normal"] = map[string]interface{}{"urls": []string{s.Data.Post.Videos.Video.Url}}
	info.downloadInfo = tmp
	return
}

func (self *ZuiYou) Obj() (obj interface{}) {
	return self
}
