package extractors

import (
	"../logs"
	"../utils"
	"fmt"
	"net/http"
	"strings"
)

type WeiBo struct {
	Base
	vid string
}

func WeiBoRegister() {
	self := new(WeiBo)
	self.Name = "weibo"
	self._VIDEO_PATTERNS = []string{`weibo\.com\/tv\/v\/(\w+)`}
	Spiders[self.Name] = self
	//self.Hd = map[string]string{"hc": "normal", "hd": "hd1", "he": "hd2", "sh": "hd3"}
	//fmt.Println(youku.Name)
}

func (self *WeiBo) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	var title string
	downloadInfo := make(map[string]interface{})
	header := make(http.Header)
	header.Add("Cookie", "SINAGLOBAL=1235634309334.4036.1466773007525; _ga=GA1.2.917481848.1468379326; __gads=ID=c7d5a0c1ec7e1367:T=1468379322:S=ALNI_Manwzqs4owCeg_MR14Ab_3x02RuOg; _s_tentry=i.youku.com; TC-V5-G0=f88ad6a0154aa03e3d2a393c93b76575; YF-Page-G0=3d55e26bde550ac7b0d32a2ad7d6fa53; YF-V5-G0=73b58b9e32dedf309da5103c77c3af4f; Apache=6927498248294.508.1494328511881; ULV=1494328511980:6:1:1:6927498248294.508.1494328511881:1481682367432; login_sid_t=782a9a2439ac71abf14466e501f2419a; YF-Ugrow-G0=ea90f703b7694b74b62d38420b5273df; SUHB=0D-0IhnQvY9mhh; UOR=www.arefly.com,widget.weibo.com,login.sina.com.cn; _T_WM=53ffe0c3d51fd89ef7843ab185b25a8a; SUB=_2AkMuRdkEdcPxrAFSkPEUymzhaYpH-jydkLDyAn7uJhMyOhh87goxqSVhUBsjZo3BcWiQbAxhynAKQZAvNA..; SUBP=0033WrSXqPxfM72wWs9jqgMF55529P9D9WFYUep4fykrDZf0mB4a5uqx5JpVsgUydNULIC44qg45qcyDdJ8oqgSQ9C4odcXt")
	pcHtml, _ := utils.GetContent(url, map[string]interface{}{"header": header})
	title = utils.R1(`info_txt\ W_f14\">([\s\S]+?)<\/div>`, pcHtml)
	title = strings.Replace(title, "\n", "", -1)
	header.Del("Cookie")
	header.Del("User-Agent")
	header.Add("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SM-N900T Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")
	mHtml, _ := utils.GetContent(url, map[string]interface{}{"header": header})
	streamUrl := utils.R1(`"stream_url":\ "(.+)"`, mHtml)
	logs.Log.Debug("streamUrl %s", streamUrl)
	info.url = url
	info.title = title
	downloadInfo["normal"] = map[string]interface{}{"urls": []string{streamUrl}}
	info.downloadInfo = downloadInfo
	return
}

func (self *WeiBo) Obj() (obj interface{}) {
	return self
}
