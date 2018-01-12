package utils

import (
	"fmt"
	"testing"
)

func resize(bar *NBar) error {
	if bar.Size < bar.Total {
		bar.Size += 10
	}
	return nil
}

func TestBar(t *testing.T) {
	bar := NewBar(100)
	bar.Resize = resize
	bar.Start()
	fmt.Println("strart")
	bar.Finish()
	fmt.Println()
	fmt.Println("finish")
}
