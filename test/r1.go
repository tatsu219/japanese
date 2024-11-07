/*
  漢字のbyte code確認用
*/

package main

import (
	"fmt"
)

func main() {
	kanji := []string{"阿", "伊", "宇"}

	for _, ctx := range kanji {
		byte_buf := []byte(ctx)
		for _, buf := range byte_buf {
			fmt.Printf("%x ", buf)
		}
		fmt.Println("")
	}

	return
}
