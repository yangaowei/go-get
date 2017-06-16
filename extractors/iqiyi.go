package extractors

import (
	"../utils"
	"fmt"
	"math/rand"
	"time"
	//"reflect"
	"crypto/md5"
	"encoding/json"
	simplejson "github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//"strings"
)

type IQiyi struct {
	Base
	Name            string
	_VIDEO_PATTERNS []string
	vid             string
	Hd              map[int]string
}

func IQiyiRegister() {
	iqiyi := new(IQiyi)
	iqiyi.Name = "iqiyi"
	iqiyi._VIDEO_PATTERNS = []string{}
	Spiders[iqiyi.Name] = iqiyi
	iqiyi.Hd = make(map[int]string)
	//{1: 'normal', 2: 'hd1', 3: 'hd2', 4: 'hd3', 5: 'hd4', 96: 'low'}
	iqiyi.Hd[1] = "normal"
	iqiyi.Hd[4] = "hd3"
	iqiyi.Hd[3] = "hd2"
	iqiyi.Hd[5] = "hd4"
	iqiyi.Hd[96] = "low"
	iqiyi.Hd[2] = "hd1"
	//fmt.Println(youku.Name)
}

func getMacid() (macid string) {
	chars := "abcdefghijklnmopqrstuvwxyz0123456789"
	size := len(chars)
	for i := 0; i < 32; i++ {
		macid += string(chars[rand.Intn(size-1)])
	}
	return macid
}

func getVf(url_params string) (vf string) {
	var sufix string
	for j := 0; j < 8; j++ {
		for k := 0; k < 4; k++ {
			v4 := 13 * (66*k + 27*j) % 35
			var v8 int
			if v4 >= 10 {
				v8 = v4 + 88
			} else {
				v8 = v4 + 49
			}
			sufix += fmt.Sprintf("%c", rune(v8))
		}
	}
	log.Println(sufix)
	url_params += sufix
	h := md5.New()
	h.Write([]byte(url_params))
	vf = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func getVps(tvid, vid string) (result []byte) {
	tm := strconv.Itoa(int(time.Now().Unix() * 1000))
	host := "http://cache.video.qiyi.com"
	src := "/vps?tvid=" + tvid + "&vid=" + vid + "&v=0&qypid=" + tvid + "_12&src=01012001010000000000&t=" + tm + "&k_tag=1&k_uid=" + getMacid() + "&rs=1"
	vf := getVf(src)
	req_url := host + src + "&vf=" + vf
	response, _ := http.Get(req_url)
	result, _ = ioutil.ReadAll(response.Body)
	return
}

func (self *IQiyi) GetVideoInfo(url string) (info VideoInfo, err error) {
	defer func() { //
		if err := recover(); err != nil {
			fmt.Println("get video info error: ", err)
		}
	}()
	html, _ := utils.GetContent(url, nil)
	tvid := utils.R1Of([]string{`data-player-tvid="([^"]+)"`, `tvid=([^&]+)`, `tvId:([^,]+)`}, html)
	vid := utils.R1Of([]string{`data-player-videoid="([^"]+)"`, `vid=([^&]+)`, `vid:([^,]+)`}, html)
	var title string
	var duration int64
	infoUrl := fmt.Sprintf("http://cache.video.qiyi.com/vi/%s/%s/", tvid, vid)
	infoVido, _ := utils.GetContent(infoUrl, nil)
	b := []byte(infoVido)
	var f interface{}
	json.Unmarshal(b, &f)
	if m, ok := (f).(map[string]interface{}); ok {
		title = (m["vn"]).(string)
		n := (m["plg"]).(float64)
		duration = int64(n)
	}
	vps := getVps(tvid, vid)
	stream, err := simplejson.NewJson(vps)
	//utils.FJson(stream)
	code, _ := stream.Get("code").String()
	streamTypes := make(map[string]interface{})
	if code == "A00000" {
		urlPrefix, _ := stream.Get("data").Get("vp").Get("du").String()
		streams, _ := stream.Get("data").Get("vp").Get("tkl").Array()
		vsArray, _ := (streams[0]).(map[string]interface{})
		vss, _ := (vsArray["vs"]).([]interface{})
		for _, vs := range vss {
			m, _ := (vs).(map[string]interface{})
			n := (m["bid"]).(json.Number)
			bid, _ := strconv.Atoi((string(n)))
			fsArray := (m["fs"]).([]interface{})
			urls := []string{}
			for _, segInfo := range fsArray {
				f, _ := (segInfo).(map[string]interface{})
				segUrl := urlPrefix + (f["l"]).(string)
				jsonData, _ := utils.GetContent(segUrl, nil)
				data, _ := simplejson.NewJson([]byte(jsonData))
				down_url, _ := data.Get("l").String()
				urls = append(urls, down_url)
			}
			tmp := make(map[string][]string)
			tmp["urls"] = urls
			streamTypes[self.Hd[bid]] = tmp
		}
	} else {
		panic("can't play this video")
	}
	info.title = title
	info.url = url
	info.duration = duration
	info.downloadInfo = streamTypes
	return
}
