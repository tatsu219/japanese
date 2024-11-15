package main

import (
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

func convertstring(arg []string, depth int, width int)[][]string{
	var result [][]string
	if width <= 0{
		err := errors.New("Invalid width value")
		return [][]string{}, err
	}else if depth <= 0{
		err := errors.New("Invalid depth value")
		return [][]string{}, err
	}
}

func make_goji(arg string) []string {
	var result []string
	var conbination [][]int
	conbination = japanese.IndexConbination([]int, 1, 3)
	for i:=0; i<len(conbination); i++{
		slice []string
		for j:=0; j<len(conbination[i]); j++{
			index := conbination[i][j] - 1
			if len(slice) == 0{

			}
		}
	}

	target := goji_array[index][0]
	target = goji_array[index][1]
	index++
	if index < len(goji_array){
		for
	}else{
		result = 
	}
	for i, ctx := range goji_array {
	}
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
			goji_buf := make_goji(mozc_format.yomi)
			for _, ctx := range goji_buf {
				result = append(result, dictionary_format{yomi: ctx, right_context_id: mozc_format.right_context_id, left_context_id: mozc_format.left_context_id, cost: mozc_format.cost, kanji: mozc_format.kanji})
			}
		}

	}

	return
}
