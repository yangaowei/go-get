package web

import (
	"../extractors"
	//"../utils"
	"../logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func test(c *gin.Context) {
	logs.Log.Informational("start:", time.Now())
	time.Sleep(5 * time.Second)
	logs.Log.Informational("end:", time.Now())
	c.String(http.StatusOK, "Hello World!")
}

func videoInfo(c *gin.Context) {
	c.String(http.StatusOK, "videoInfo")
	url := c.DefaultQuery("url", "mis")
	if url == "mis" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "mis params url",
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
		logs.Log.Informational("get IE %s", key)
		if len(key) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "暂不支持该站点",
			})
		} else {
			info, err := spider.GetVideoInfo(url)
			if err == nil {
				c.JSON(http.StatusOK, info.Dumps())
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": fmt.Sprintf("%v", err),
				})
			}
		}
	}
}

func Run() {
	router := gin.Default()
	router.GET("/", test)
	router.GET("/video/info", videoInfo)
	router.StaticFile("/favicon.ico", "./web/resources/favicon.ico")
	router.Run(":8001")
}
