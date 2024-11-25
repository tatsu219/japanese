package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	original_dictionary_dir   = "../../original-dictionary/"
	original_dictionary_name  = "dictionary"
	original_dictionary_count = 10
	mozc_dictionary_dir       = "../../mozc-dictionary/"
	mozc_dictionary_name      = "dictionary"
)

type mozc_struct struct {
	file             string
	no               int
	yomi             string
	left_context_id  string
	right_context_id string
	cost             string
	kanji            string
}

var dictionary map[string][]mozc_struct

func main() {
	dictionary = make(map[string][]mozc_struct)
	for i := range original_dictionary_count {
		file_num := fmt.Sprintf("%02d", i)
		filename := original_dictionary_dir + original_dictionary_name + file_num + ".txt"
		fmt.Println(filename)
		fd, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("os.Open error")
			return
		}
		defer fd.Close()

		input, err := io.ReadAll(fd)
		if err != nil {
			fmt.Println(err)
			fmt.Println("os.Open error")
			return
		}

		slice := strings.Split(string(input), "\n")
		if err != nil {
			fmt.Println(err)
			fmt.Println("strings.Split error")
			return
		}

		if slice[len(slice)-1] == "" {
			slice = slice[:len(slice)-1]
		}

		for i := 0; i < len(slice); i++ {
			buf := strings.Split(slice[i], "\t")
			dictionary[buf[0]] = append(dictionary[buf[0]], mozc_struct{file: filename, no: i + 1, yomi: buf[0], left_context_id: buf[1], right_context_id: buf[2], cost: buf[3], kanji: buf[4]})
		}
	}

	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("<")
		input.Scan()
		str := input.Text()

		result := dictionary[str]
		if result == nil {
			fmt.Println("no match")
			continue
		}
		for _, ctx := range result {
			fmt.Println(ctx)
		}
	}
}
