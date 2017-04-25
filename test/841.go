package main

import (
	"fmt"
	"time"
)

func append(ch chan int, i int) {
	fmt.Println(i)
	ch <- i
	//fmt.Println(i)
}

func main() {
	ch := make(chan int, 3)

	for i := 0; i < 10; i++ {
		go append(ch, i)
	}
	time.Sleep(10 * time.Second)
}
