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
	output_file     = "additional.txt"
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

func convertstring(arg string, depth int, width int) ([]string, error) {
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
	if err != nil {
		return []string{}, err
	}
	if (depth + 1) <= width {
		return_value, err := convertstring(buf, depth+1, width)
		if err != nil {
			return []string{}, err
		}
		result = append(result, return_value...)
	} else {
		result = append(result, buf)
	}
	for i := 0; i < len(conbination); i++ {
		return_value := japanese.StrconvSelect(buf, goji_array[depth-1][1], goji_array[depth-1][0], conbination[i])
		if (depth + 1) <= width {
			slice, err := convertstring(return_value, depth+1, width)
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
	result, err := convertstring(arg, 1, len(goji_array))
	return result, err
}

func main() {
	file_index := 0
	var result []dictionary_format
	for file_index < 10 {
		filename := fmt.Sprintf("%02d", file_index)
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

	fd, err := os.Create(output_dir + output_file)
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Create additional.txt error")
		return
	}
	defer fd.Close()

	for i, ctx := range result {
		buf := ctx.yomi + "\t" + strconv.Itoa(ctx.right_context_id) + "\t" + strconv.Itoa(ctx.left_context_id) + "\t" + strconv.Itoa(ctx.cost) + "\t" + ctx.kanji
		fd.Write([]byte(buf))
		if i != len(result)-1 {
			buf += "\n"
		}
	}

	return
}
