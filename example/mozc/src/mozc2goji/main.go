package main

import (
	"errors"
	"fmt"
	"github.com/hasuburero/japanese/japanese"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	dictionary_path = "../../mozc-dictionary/"
	dictionary_name = "dictionary"
	output_dir      = "../../additional-dictionary/"
	output_file     = "additional2"
	dic_range       = 10
)

const (
	zu = "ず"
	du = "づ"
	o  = "お"
	u  = "う"
	ji = "じ"
	di = "ぢ"
)

var goji_array [][]string = [][]string{{zu, du}, {o, u}, {ji, di}}

type dictionary_format struct {
	yomi             string
	right_context_id int
	left_context_id  int
	cost             int
	kanji            string
}

func convertstring(arg string, original string, depth int, width int) ([]string, error) {
	var result []string
	if width <= 0 {
		err := errors.New("Invalid width value")
		return []string{}, err
	} else if depth <= 0 {
		err := errors.New("Invalid depth value")
		return []string{}, err
	}
	buf := japanese.StrconvAll(arg, goji_array[depth-1][0], goji_array[depth-1][1])
	goji_count := japanese.Strcount(buf, goji_array[depth-1][1])
	conbination, err := japanese.IndexConbination([]int{}, 1, goji_count)
	if err == errors.New("Invalid depth error") {
		return []string{}, err
	}
	if (depth + 1) <= width {
		return_value, err := convertstring(buf, original, depth+1, width)
		if err != nil {
			return []string{}, err
		}
		result = append(result, return_value...)
	} else {
		if buf == original {
			return []string{}, nil
		}
		result = append(result, buf)
	}
	for i := 0; i < len(conbination); i++ {
		return_value := japanese.StrconvSelect(buf, goji_array[depth-1][1], goji_array[depth-1][0], conbination[i])
		if (depth + 1) <= width {
			slice, err := convertstring(return_value, original, depth+1, width)
			if err != nil {
				return []string{}, err
			}
			result = append(result, slice...)
		} else {
			result = append(result, return_value)
		}
	}

	return result, nil
}

func make_goji(arg string) ([]string, error) {
	result, err := convertstring(arg, arg, 1, len(goji_array))
	for i, ctx := range result {
		if ctx == arg {
			if i == 0 && (i+1) < len(result) {
				result = result[1:]
			} else if (i + 1) == len(result) {
				result = []string{}
			} else if i != len(result)-1 {
				result = append(result[:i], result[i+1:]...)
			} else {
				result = result[:i]
			}
		}
	}
	return result, err
}

func main() {
	file_index := 0
	var result []dictionary_format
	for file_index < dic_range {
		filename := fmt.Sprintf("%02d", file_index)
		fmt.Println(filename + ".txt")
		file_index++
		filename = dictionary_path + dictionary_name + filename + ".txt"
		fd, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("os.Open error")
			fmt.Printf("filename: %s\n", filename)
			return
		}
		defer fd.Close()

		buf, err := io.ReadAll(fd)
		if err != nil {
			fmt.Println(err)
			fmt.Println("io.ReadAll error")
			return
		}

		slice := strings.Split(string(buf), "\n")
		if slice[len(slice)-1] == "" {
			slice = slice[:len(slice)-1]
		}

		for _, ctx := range slice {
			buf := strings.Split(ctx, "\t")
			right_id, err := strconv.Atoi(buf[1])
			if err != nil {
				fmt.Println(err)
				fmt.Println("strconv.Atoi error")
				return
			}
			left_id, err := strconv.Atoi(buf[2])
			if err != nil {
				fmt.Println(err)
				fmt.Println("strconv.Atoi error")
				return
			}
			cost, err := strconv.Atoi(buf[3])
			if err != nil {
				fmt.Println(err)
				fmt.Println("strconv.Atoi error")
				return
			}

			mozc_format := dictionary_format{yomi: buf[0], right_context_id: right_id, left_context_id: left_id, cost: cost, kanji: buf[4]}
			goji_buf, err := make_goji(mozc_format.yomi)
			if err != nil {
				fmt.Println(err)
				fmt.Println("make_goji error")
				return
			}
			for _, ctx := range goji_buf {
				result = append(result, dictionary_format{yomi: ctx, right_context_id: mozc_format.right_context_id, left_context_id: mozc_format.left_context_id, cost: mozc_format.cost, kanji: mozc_format.kanji})
			}
		}
	}

	index_start := 0
	index_end := 0
	result_length := len(result)
	if result_length < 100000 {
		index_end = result_length
	} else {
		index_end = 100000
	}
	for file_index := 0; ; file_index++ {
		file_num := fmt.Sprintf("%02d", file_index)
		filename := output_dir + output_file + file_num + ".txt"
		fmt.Println(filename)
		fd, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("os.Create error")
			return
		}
		defer fd.Close()
		result_buf := result[index_start:index_end]
		for i, ctx := range result_buf {
			buf := ctx.yomi + "\t" + strconv.Itoa(ctx.right_context_id) + "\t" + strconv.Itoa(ctx.left_context_id) + "\t" + strconv.Itoa(ctx.cost) + "\t" + ctx.kanji
			if i != len(result_buf)-1 {
				buf += "\n"
			}
			fd.Write([]byte(buf))
		}
		if index_end == result_length {
			break
		}
		if index_end < result_length {
			index_start = index_end
			if index_end+100000 < result_length {
				index_end += 100000
			} else {
				index_end = result_length
			}
		}
	}

	return
}
