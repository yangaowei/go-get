package utils

import (
	"fmt"
	"strconv"
	"time"
)

type NBar struct {
	Total  int
	Size   int
	finish bool
	start  time.Time
	//end    time.Time
	Resize func(bar *NBar) error
}

func NewBar(Total int) *NBar {
	return &NBar{Total: Total, Size: 0}
}

func (bar *NBar) Start() {
	bar.start = time.Now()
	go func() {
		for {
			if bar.Size < bar.Total {
				bar.Resize(bar)
				time.Sleep(1 * time.Second)
			} else {
				bar.finish = true
				return
			}
		}
	}()
}

func (bar *NBar) Finish() {
	for !bar.finish {
		bar.print()
		time.Sleep(1 * time.Second)
	}
	bar.print()
}

func (bar *NBar) cost() string {
	s := fmt.Sprintf("%ds", time.Now().Second()-bar.start.Second())
	return s
}

func (bar *NBar) print() {
	str := ""
	size := 100
	//fmt.Println(bar.Size)
	count := bar.Size * size / bar.Total
	//fmt.Println(count)
	for i := 0; i < size; i++ {
		if i < count {
			str += "="
		} else {
			str += "_"
		}
	}
	//str = "[" + str + "] " + strconv.Itoa(count) + "%" + "  "
	str = fmt.Sprintf("[%s] %s%  cost %s", str, strconv.Itoa(count), bar.cost())
	fmt.Printf("\r%s", str)
}
