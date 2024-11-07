package utf8

import ()

const (
	bit_shift_width = 8
	mask1byte       = 0xff000000
	mask2byte       = 0xffff0000
	mask3byte       = 0xffffff00
	maskflip        = 0xffffffff
)

func Int2byte(arg int) []byte {
	var byte_count int = 0
	var buf []byte
	switch {
	case arg&mask1byte != 0:
		byte_count = 4
	case arg&mask2byte != 0:
		byte_count = 3
	case arg&mask3byte != 0:
		byte_count = 2
	default:
		byte_count = 1
	}

	for i := range byte_count {
		switch byte_count - i {
		case 4:
			buf = append(buf, byte(arg>>(8*(4-1))))
		case 3:
			buf = append(buf, byte((arg&(mask1byte^maskflip))>>8*(3-1)))
		case 2:
			buf = append(buf, byte((arg&(mask2byte^maskflip))>>8*(2-1)))
		case 1:
			buf = append(buf, byte((arg&(mask3byte^maskflip))>>8*(1-1)))
		}
	}

	return buf
}

func Byte2int(arg []byte) int {
	var shift_width int = len(arg)
	var buf int = 0x00000000
	for i := range shift_width {
		buf += int(arg[i]) << (bit_shift_width * (shift_width - 1 - i))
	}

	return buf
}
