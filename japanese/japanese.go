package japanese

import (
	"errors"
	"fmt"
	"strings"
)

const (
	byte4 = 0xf0
	byte3 = 0xe0
	byte2 = 0xc2
	byte1 = 0x00
)

func swap(arg1, arg2 int) (int, int) {
	return arg2, arg1
}

func RuneSubstring(arg string, a, b int) (string, error) {
	var result string = ""
	var index int = 0
	arg_length := RuneLength(arg)
	if a < 0 {
		return "", errors.New("invalid start index")
	} else if a == b {
		return "", errors.New("invalid stop index")
	} else if b > arg_length {
		return "", errors.New("out of size arg")
	}
	for _, ctx := range arg {
		if index == b {
			break
		} else if index >= a {
			result += string(ctx)
		}
		index++
	}

	return result, nil
}

func RuneLength(arg string) int {
	count := 0
	for _, _ = range arg {
		count++
	}
	return count
}

func Sort(arg []int) []int {
	length := len(arg)
	if length < 2 {
		return arg
	}
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if arg[j] > arg[j+1] {
				arg[j], arg[j+1] = swap(arg[j], arg[j+1])
			}
		}
	}
	return arg
}

func IndexConbination(arg []int, depth int, width int) ([][]int, error) {
	var result [][]int
	if width <= 0 {
		err := errors.New("Invalid width value")
		return [][]int{}, err
	} else if depth <= 0 {
		err := errors.New("Invalid depth value")
		return [][]int{}, err
	}
	if depth == 1 {
		for i := 1; i <= width; i++ {
			buf := []int{i}
			result = append(result, buf)
			if (depth + 1) <= width {
				return_value, err := IndexConbination(buf, depth+1, width)
				if err != nil {
					return [][]int{}, err
				}
				result = append(result, return_value...)
			}
		}
	} else {
		for i := arg[depth-2] + 1; i <= width; i++ {
			buf := append(arg, i)
			result = append(result, buf)
			if (depth + 1) <= width {
				return_value, err := IndexConbination(buf, depth+1, width)
				if err != nil {
					return [][]int{}, err
				}
				result = append(result, return_value...)
			}
		}
	}
	return result, nil
}

func Strcount(arg string, target string) int {
	slice := strings.Split(arg, target)
	return len(slice) - 1
}

func StrconvFirst(arg string, target string, dest string) string {
	slice := strings.Split(arg, target)
	length := len(slice)
	if length < 2 {
		return arg
	}
	result := ""
	for i := 1; i < length; i++ {
		if i == 1 {
			result += dest + slice[i]
		} else {
			result += target + slice[i]
		}
	}
	return result
}

func StrconvLast(arg string, target string, dest string) string {
	slice := strings.Split(arg, target)
	length := len(slice)
	if length < 2 {
		return arg
	}

	result := ""
	for i := 1; i < length; i++ {
		if i == (length - 1) {
			result += dest + slice[length-1]
		} else {
			result += target + slice[i]
		}
	}
	return result
}

func StrconvAll(arg string, target string, dest string) string {
	slice := strings.Split(arg, target)
	length := len(slice)
	if length < 2 {
		return arg
	}
	result := slice[0]
	for i := 1; i < length; i++ {
		result += dest + slice[i]
	}

	return result
}

func StrconvSelect(arg string, target string, dest string, index []int) string {
	for i, ctx := range index {
		if ctx <= 0 {
			if i != 0 {
				if i != len(index)-1 {
					index = append(index[:i], index[i+1:]...)
				} else {
					index = index[:i]
				}
			} else {
				if i != len(index)-1 {
					index = index[i+1:]
				} else {
					index = []int{}
					return arg
				}
			}
		}
	}
	index = Sort(index)
	slice := strings.Split(arg, target)
	length := len(slice)
	length_index := len(index)
	if length < 2 {
		return arg
	}

	result := slice[0]
	for i, j := 1, 0; i < length; i++ {
		conv_flag := false
		if j < length_index {
			if i == index[j] {
				result += dest + slice[i]
				conv_flag = true
				j++
			}
		}
		if !conv_flag {
			result += target + slice[i]
		}
	}

	return result
}

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
