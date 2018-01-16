package utils

import (
	"fmt"
	"testing"
)

func resize(bar *NBar) error {
	if bar.Size < bar.Total {
		bar.Size += int64(70000)
		if bar.Size > bar.Total {
			bar.Size = bar.Total
		}
	}
	return nil
}

func TestBar(t *testing.T) {
	var total int64
	total = 7116238
	bar := NewBar(total)
	bar.Resize = resize
	bar.Start()
	bar.Finish()
	fmt.Println()
}
