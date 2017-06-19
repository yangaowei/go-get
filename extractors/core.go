package extractors

import (
	"../utils"
	"github.com/PuerkitoBio/goquery"
	simplejson "github.com/bitly/go-simplejson"
	"log"
	"strings"
	"time"
)

type VideoInfo struct {
	title        string
	url          string
	duration     int64
	downloadInfo map[string]interface{}
}

type Core interface {
	GetVideoInfo(url string) (info VideoInfo, err error)
	GetHtml(url string) (html string, err error)
	Obj() (obj interface{})
}

func (self *VideoInfo) dumps() (info map[string]interface{}) {
	info = make(map[string]interface{})
	info["title"] = self.title
	info["url"] = self.url
	info["duration"] = self.duration
	info["downloadInfo"] = self.downloadInfo
	return info
}

//实例基类
type Base struct {
}

func (base *Base) CurrentTime() (ts int64) {
	return time.Now().Unix()
}

func (base *Base) GetVideoInfo(url string) (info VideoInfo, err error) {
	return VideoInfo{}, nil
}

func (base *Base) GetHtml(url string) (html string, err error) {
	log.Println("request url ", url)
	return url + "html", nil
}

func (base *Base) Obj() (obj interface{}) {
	return base
}

func (base *Base) BuildDoc(url string) (doc *goquery.Document, err error) {
	log.Println("build doc ", url)
	doc, err = goquery.NewDocument(url)
	return
}

func (base *Base) BuildJson(url string) (json *simplejson.Json, err error) {
	video_html, err := utils.GetContent(url, nil)
	bjson := []byte(video_html)
	json, err = simplejson.NewJson(bjson)
	return
}

var (
	Spiders = make(map[string]Core)
)

func init() {
	YouKuRegister()
	QQRegister()
	IQiyiRegister()
	SohuRegister()
}

func GetExtractor(url string) (extractor Core) {
	log.Println(url)
	if strings.Contains(url, "youku") {
		return Spiders["youku"]
	}
	return nil
}
