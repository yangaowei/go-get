package extractors

import (
	"../logs"
	"../utils"
	"encoding/xml"
	"fmt"
	"strconv"
)

type PPTV struct {
	Base
	vid string
}

type XMLResult struct {
	Channel struct {
		Dur       string `xml:"dur,attr"`
		Timestamp string `xml:"timestamp,attr"`
		Titile    string `xml:"nm,attr"`
		File      struct {
			Item []struct {
				Rid string `xml:"rid,attr"`
				Ft  string `xml:"ft,attr"`
			} `xml:"item"`
		} `xml:"file"`
	} `xml:"channel"`
	Dt []struct {
		Ft  string `xml:"ft,attr"`
		Sh  string `xml:"sh"`
		Key string `xml:"key"`
	} `xml:"dt"`

	Dragdata []struct {
		Ft  string `xml:"ft,attr"`
		Sgm []struct {
			No  string `xml:"no,attr"`
			Rid string `xml:"rid,attr"`
			Hl  string `xml:"rid,hl"`
		} `xml:"sgm"`
	} `xml:"dragdata"`
}

func parseXML(input string) (result *XMLResult) {
	result = &XMLResult{}
	err := xml.Unmarshal([]byte(input), result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	//fmt.Println(v)
	return result
}

func PPTVRegister() {
	self := new(PPTV)
	self.Name = "pptv"
	self._VIDEO_PATTERNS = []string{`v\.pptv\.com\/show\/(\w+)\.html`}
	Spiders[self.Name] = self
	self.Hd = make(map[string]string)
	self.Hd["4"] = "hd4"
	self.Hd["3"] = "hd3"
	self.Hd["2"] = "hd2"
	self.Hd["1"] = "hd1"
	self.Hd["0"] = "normal"
	//fmt.Println(youku.Name)
}

func (self *PPTV) GetVid(url string) (vid string) {
	if doc, err := self.BuildDoc(url); err == nil {
		fmt.Println(doc)
	}
	return
}

func (self *PPTV) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	logs.Log.Debug("len %d", len(html))
	vid := utils.R1(`webcfg\s*=\s*{"id":\s*(\d+)`, html)
	logs.Log.Debug("vid %s", vid)
	if len(vid) == 0 {
		panic("vid is null")
	}
	VideoInfoUrl := fmt.Sprintf("http://web-play.pptv.com/webplay3-0-%s.xml?type=web.fpp&version=4", vid)
	xml, _ := utils.GetContent(VideoInfoUrl, nil)
	//logs.Log.Debug("XML %s", xml)

	xmlResult := parseXML(xml)
	stringsHd := make(map[string]interface{})
	for _, item := range xmlResult.Channel.File.Item {
		fileName := item.Rid
		ft := item.Ft
		var key string
		var serverHost string
		for _, dt := range xmlResult.Dt {
			if dt.Ft == ft {
				key = dt.Key
				serverHost = dt.Sh
				break
			}
		}
		cdnUrlByFormat := "http://%s/%d/%s/0/%s?fpp.ver=1.3.0.19&type=web.fpp&k=%s"
		var urls []string
		for _, drag := range xmlResult.Dragdata {
			if drag.Ft == ft {
				for index, sgm := range drag.Sgm {
					urls = append(urls, fmt.Sprintf(cdnUrlByFormat, serverHost, index, sgm.Hl, fileName, key))
				}
				break
			}
		}
		stringsHd[self.Hd[ft]] = map[string]interface{}{"urls": urls}
	}
	//fmt.Println(xmlResult)
	//var title string
	// var duration int64
	// var createTime int64

	info.title = xmlResult.Channel.Titile
	info.url = url
	info.duration, _ = strconv.ParseInt(xmlResult.Channel.Dur, 10, 64)
	info.downloadInfo = stringsHd //20170410100731
	logs.Log.Debug("Timestamp %s", xmlResult.Channel.Timestamp)
	info.createTime = utils.StringToMilliseconds("20060102150405", xmlResult.Channel.Timestamp)

	return
}

func (self *PPTV) Obj() (obj interface{}) {
	return self
}
