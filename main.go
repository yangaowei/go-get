package main

import (
	"./extractors"
	"./utils/surfer"
	"fmt"
	"io/ioutil"
)

func main() {
	core := extractors.GetExtractor("youku")
	info, _ := core.GetVideoInfo("youku")
	fmt.Println(info)
	fmt.Println(core.GetHtml("test"))

	req := &surfer.DefaultRequest{Url: "https://www.baidu.com"}

	fmt.Println(req.GetUrl())
	fmt.Println(req.GetHeader())
	download := surfer.New()
	resp, _ := download.Download(req)
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	resp.Body.Close()
	fmt.Println(req.GetHeader())
}
