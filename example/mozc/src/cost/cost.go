package main

import (
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

func main() {
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

		for i, ctx := range slice {
			buf := strings.Split(ctx, "\t")
			if buf[3] == "0" {
				fmt.Println(i, buf)
			}
		}
		fmt.Println("")
	}

	return
}
