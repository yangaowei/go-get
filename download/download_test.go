package download

import (
	"log"
	//"os"
	"testing"
)

func TestUrlSave(t *testing.T) {
	result := UrlSave("test.mp4", "http://vssauth.waqu.com/baidutieba/05toxsv7qddqd781.mp4?auth_key=1498227820-0-0-312221343a0a5c32cbcfcc8b1c6b253c")
	log.Println(result)
}
