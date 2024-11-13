package utf8

import (
	"fmt"
)

const (
	byte4 = 0xf0
	byte3 = 0xe0
	byte2 = 0xc2
	byte1 = 0x00
)

func Printbyte(arg []byte) {
	for _, ctx := range arg {
		fmt.Printf("%x ", ctx)
	}
	fmt.Println("")
}

func Splitutf8(arg string) [][]byte {
	arg_byte := []byte(arg)
	var result [][]byte
	for i := 0; i < len(arg_byte); {
		inc_value := 0
		switch {
		case arg_byte[i] >= byte4:
			inc_value = 4
		case arg_byte[i] >= byte3:
			inc_value = 3
		case arg_byte[i] >= byte2:
			inc_value = 2
		default:
			inc_value = 1
		}

		result = append(result, arg_byte[i:i+inc_value])
		i += inc_value
	}

	return result
}
