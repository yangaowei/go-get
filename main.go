package main

//
import (
	"./download"
	"./extractors"
	"./logs"
	"./web"
	//"encoding/json"
	"flag"
	"os"
)

var (
	port    string
	pattern string //api  cmd
	path    string //api  cmd
	//url     string
)

func initFlag() {
	flag.StringVar(&port, "port", "8002", "server port")
	flag.StringVar(&pattern, "p", "cmd", "runing pattern")
	//flag.StringVar(&path, "path", ".", "download path")
	//flag.StringVar(&url, "u", ".", "download path")
	flag.Parse()
}

func main() {
	initFlag()
	//
	logs.Log.Informational("pattern: %s", pattern)
	if pattern == "api" {
		web.Run(port)
	} else {
		url := os.Args[1]
		logs.Log.Informational("url: %s", url)
		key, spider := extractors.GetExtractor(url)
		logs.Log.Informational("get IE %s", key)
		if len(key) == 0 {
			logs.Log.Warning("暂不支持该站点")
		} else {
			i, err := spider.GetVideoInfo(url)
			if err == nil {
				info := i.Dumps()
				//sjson, err := json.MarshalIndent(info, "", "\t")
				info["site"] = key
				dowloadInfo := info["downloadInfo"].(map[string]interface{})
				for _, hd := range []string{"hd3", "hd2", "hd1", "normal"} {
					if _, ok := dowloadInfo[hd]; ok {
						urls := i.Urls(hd)
						//logs.Log.Warning("urls %v", urls)
						info["type"] = hd
						download.Download(urls, "mp4", info)
						break
					}
				}

			} else {
				logs.Log.Warning("解析失败 %v", err)
			}
		}
	}
}
