package utils

import (
	"fmt"
	"testing"
)

func resize(bar *NBar) error {
	if bar.Size < bar.Total {
		bar.Size += 1
	}
	return nil
}

func TestBar(t *testing.T) {
	bar := NewBar(100)
	bar.Resize = resize
	bar.Start()
	bar.Finish()
	fmt.Println()
}
