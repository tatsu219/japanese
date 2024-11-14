/*
平仮名片仮名 converter test program
*/

package main

import (
	"fmt"
	"github.com/hasuburero/japanese/japanese"
)

func main() {
	var test string = "あかさたなハマヤラワ"
	result := japanese.ConvHiraKata(test)
	fmt.Println(result)
	for _, ctx := range []byte(result) {
		fmt.Printf("%x ", ctx)
	}
	fmt.Println("")
}
