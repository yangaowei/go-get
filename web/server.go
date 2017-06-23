package web

import (
	"../extractors"
	//"../utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func test(c *gin.Context) {
	log.Println("start:", time.Now())
	time.Sleep(5 * time.Second)
	log.Println("end:", time.Now())
	c.String(http.StatusOK, "Hello World!")
}

func videoInfo(c *gin.Context) {
	c.String(http.StatusOK, "videoInfo")
	url := c.DefaultQuery("url", "mis")
	if url == "mis" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "mis url",
		})
	} else {
		var key string
		var spider extractors.Core
		for a, b := range extractors.Spiders {
			if b.MatchUrl(url) {
				key = a
				spider = b
				break
			}
		}
		log.Println("get IE ", key)
		if len(key) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "暂不支持该站点",
			})
		} else {
			info, _ := spider.GetVideoInfo(url)
			c.JSON(http.StatusOK, info.Dumps())
		}
	}
}

func Run() {

	router := gin.Default()
	router.GET("/", test)
	router.GET("/video/info", videoInfo)
	router.Run(":8001")
}
