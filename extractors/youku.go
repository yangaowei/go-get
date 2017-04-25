package extractors

import (
	"fmt"
)

type YouKu struct {
	Base
	Name string
}

func YouKuRegister() {
	youku := new(YouKu)
	youku.Name = "youku"
	Spiders[youku.Name] = youku
}

func (obj *YouKu) GetVideoInfo(url string) (info VideoInfo, err error) {
	return VideoInfo{url: url}, nil
}
