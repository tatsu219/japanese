package main

import (
	"fmt"
	"github.com/hasuburero/japanese/japanese"
)

func main() {
	input := "ぐらたんはぐらたん"
	target := "ぐ"
	dest := "が"
	result := japanese.StrconvAll(input, target, dest)
	fmt.Println(result)
	result = japanese.StrconvFirst(input, target, dest)
	fmt.Println(result)
	result = japanese.StrconvLast(input, target, dest)
	fmt.Println(result)
	result = japanese.StrconvSelect(input, target, dest, []int{2})
	fmt.Println(result)
	result = japanese.StrconvSelect(input, target, dest, []int{1, 2})
	fmt.Println(result)
	result = japanese.StrconvSelect(input, target, dest, []int{1})
	fmt.Println(result)
	result = japanese.StrconvSelect(input, target, dest, []int{0})
	fmt.Println(result)

	target = "あ"
	dest = "い"
	result = japanese.StrconvAll(input, target, dest)
	fmt.Println(result)
	return
}
