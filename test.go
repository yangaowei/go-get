package main

import (
	//"./extractors"
	"fmt"
)

type BI interface {
	Less(s string)
}

type Base struct {
	Name string
}

type Student struct {
	Base
	Score int64
}

type Core interface {
	GetVideoInfo(url string)
}

//Spiders := make(map[string]Core)
//b := make(map[string]Core)
var (
	Spiders = make(map[string]Core)
)

func (base Base) GetVideoInfo(url string) {
	fmt.Println(url)
}

func main() {
	// var base BI
	// fmt.Println(base)
	// base = Base{"test"}
	// base.less()
	Spiders["test"] = Student{}
	fmt.Println(Spiders)
	Spiders["test"].GetVideoInfo("url")
}
