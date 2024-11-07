/*
平仮名片仮名 converter test program
*/

package main

import (
	"fmt"
	"japanese/hirakata"
)

func main() {
	var test string = "あかさたなハマヤラワ"
	result := hirakata.ConvHiraKata(test)
	fmt.Println(result)
	for _, ctx := range []byte(result) {
		fmt.Printf("%x ", ctx)
	}
	fmt.Println("")
}
