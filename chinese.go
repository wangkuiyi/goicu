package main

import (
	"fmt"
)

func StrictlyChinese(text string) error {
	for i, c := range text {
		if c < 0x4e00 || 0x9fcc < c {
			return fmt.Errorf("The %d-th rune %c(%#x) is not Chinese", i, c, c)
		}
	}
	return nil
}
