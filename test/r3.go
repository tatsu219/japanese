package main

import (
	"fmt"
	"github.com/hasuburero/japanese/utf8"
	"strings"
)

func main() {
	input := "ぐらたんはぐらたん"
	split := "ぐ"
	buf := strings.Split(input, split)
	if len(buf) == 1 {
		fmt.Printf("target does not include %s\n", split)
		return
	}

	fmt.Println(buf)
	return
}
