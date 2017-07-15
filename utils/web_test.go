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

func TestFindAll(t *testing.T) {
	content := FindAll("abcd([0-9]+)", "abcd1234asdfadfdadf123123afaf1212")
	log.Println("TestFindAll:", content)
}

func TestFindSubAll(t *testing.T) {
	content := FindSubAll("abcd([0-9]+)", "abcd1234asdfadfdadf123123afaf1212")
	log.Println("FindSubAll:", content)
}

func TestR1(t *testing.T) {
	content := R1("abcddd([0-9]+)", "time: '2017/06/07',")
	log.Println("TestR1:", content, "len:", len(content))
}

func TestR1of(t *testing.T) {
	patterns := []string{"abcd([0-9]+)"}
	content := R1Of(patterns, "abcd1234asdfadfdadf123123afaf1212")
	log.Println("TestR1of:", content, "len:", len(content))
}

func TestLoads(t *testing.T) {
	result := Loads(`{"a":1}`)
	log.Println(result)
	log.Println(result["a"])
}
