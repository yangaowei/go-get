package utils

import (
	"fmt"
	"github.com/cnych/starjazz/mathx"
	"strconv"
	"time"
)

type NBar struct {
	Total  int64
	Size   int64
	finish chan bool
	start  int64
	//end    time.Time
	Resize func(bar *NBar) error
}

func NewBar(Total int64) *NBar {
	return &NBar{Total: Total, finish: make(chan bool, 1)}
}

func (bar *NBar) Start() {
	bar.start = time.Now().Unix()
	go func() {
		for {
			if bar.Size < bar.Total {
				bar.Resize(bar)
				time.Sleep(100 * time.Millisecond)
			} else {
				bar.finish <- true
				return
			}
		}
	}()
}

func (bar *NBar) Finish() {
	t := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-bar.finish:
			goto ForEnd
		case <-t:
			bar.print()
		}
	}
ForEnd:
	bar.print()
}

func (bar *NBar) cost() string {
	s := fmt.Sprintf(" %ds", time.Now().Unix()-bar.start)
	return s
}

func formatSize(size int64) string {
	s := fmt.Sprintf("%.2f MiB", mathx.Round(float64(size)/1024/1024, 2))
	return s
}

func (bar *NBar) print() {
	str := ""
	var size int64
	size = 100
	//fmt.Println(bar.Size)
	count := bar.Size * size / bar.Total
	//fmt.Println(count)
	for i := int64(0); i < size; i++ {
		if i < count {
			str += "="
		} else if i == count {
			str += ">"
		} else {
			str += "_"
		}
	}
	//str = "[" + str + "] " + strconv.Itoa(count) + "%" + "  "
	str = fmt.Sprintf("%s/%s [%s] %s%% %s", formatSize(bar.Size), formatSize(bar.Total), str, strconv.FormatInt(count, 10), bar.cost())
	fmt.Printf("\r%s", str)
}
