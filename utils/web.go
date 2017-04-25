package utils

import (
	"./surfer"
	"io/ioutil"
	"log"
)

func GetHtml(req surfer.Request) (resp string, err error) {
	log.Println("get html from url ", req.GetUrl())
	down := surfer.New()
	if response, e := down.Download(req); e == nil {
		bytes, _ := ioutil.ReadAll(response.Body)
		resp = string(bytes)
		response.Body.Close()
	} else {
		log.Println("err", e)
		resp = "resp"
		err = e
	}
	return resp, err
}
