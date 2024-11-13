package hirakata

import (
	"../utf8"
	"fmt"
)

const (
	hira_start  = 0xe38180
	hira_middle = 0xe38280
	hira_end    = 0xe382a0
	kata_start  = 0xe382a0
	kata_middle = 0xe38380
	kata_end    = 0xe383c0

	shift_width_range  = 0x000020
	top_shift_width    = kata_start - hira_start
	middle_shift_width = kata_middle - (hira_start + shift_width_range)
	bottom_shift_width = (kata_middle + shift_width_range) - hira_middle
)

func ConvHiraKata(arg string) string {
	byte_args := utf8.Splitutf8(arg)
	var result string = ""
	for _, ctx := range byte_args {
		var buf []byte
		int_buf := utf8.Byte2int(ctx)
		if hira_start <= int_buf && int_buf < kata_end {
			switch {
			case int_buf < hira_start+shift_width_range:
				int_buf += top_shift_width
			case int_buf < hira_start+shift_width_range*2:
				int_buf += middle_shift_width
			case int_buf < hira_end:
				int_buf += bottom_shift_width
			case int_buf < kata_start+shift_width_range:
				int_buf -= top_shift_width
			case int_buf < kata_middle+shift_width_range:
				int_buf -= middle_shift_width
			case int_buf < kata_end:
				int_buf -= bottom_shift_width
			}
			buf = utf8.Int2byte(int_buf)
			fmt.Printf("%x\n", int_buf)
			utf8.Printbyte(buf)
			fmt.Printf("%s\n", string(buf))
		} else {
			buf = ctx
		}
		result += string(buf)
	}

	return result
}
