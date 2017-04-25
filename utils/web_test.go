package utils

import (
	"./surfer"
	"github.com/PuerkitoBio/goquery"
	"log"
	//"reflect"
	"testing"
)

func TestGetHtml(t *testing.T) {
	request := &surfer.DefaultRequest{Url: "http://www.waqu.com", TryTimes: 1}
	request.GetUrl()
	resp, err := GetHtml(request)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(resp))
}

func TestQuery(t *testing.T) {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		info := contentSelection.Find(".title a")
		url, _ := info.Attr("href")
		log.Println("第", i+1, "个帖子的标题：", info.Text(), "url:", url)
	})

}
