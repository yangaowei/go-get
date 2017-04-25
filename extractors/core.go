package extractors

import (
	"log"
	"strings"
)

type VideoInfo struct {
	title        string
	url          string
	duration     int64
	downloadInfo map[string]string
}

type Core interface {
	GetVideoInfo(url string) (info VideoInfo, err error)
	GetHtml(url string) (html string, err error)
}

//实例基类
type Base struct {
}

func (base *Base) GetVideoInfo(url string) (info VideoInfo, err error) {
	return VideoInfo{}, nil
}

func (base *Base) GetHtml(url string) (html string, err error) {
	log.Println("request url ", url)
	return url + "html", nil
}

var (
	Spiders = make(map[string]Core)
)

func init() {
	YouKuRegister()
}

func GetExtractor(url string) (extractor Core) {
	log.Println(url)
	if strings.Contains(url, "youku") {
		return Spiders["youku"]
	}
	return nil
}
