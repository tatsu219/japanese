package main

import (
	"fmt"
	"strconv"
	"github.com/hasuburero/japanese"
	"io"
	"os"
)

const (
	dictionary_path = "../../mozc-dictionary/"
	dictionary_name = "dictionary0"
)

func main() {
	file_index := 0
	for {
		filename := strconv.Itoa(file_index++)
		filename = dictionary_path + filename
		fd, err := os.Open(filename)
		if err != nil{
			fmt.Println(err)
			fmt.Println("os.Open error")
			fmt.Printf("filename: %s\n", filename)
			return
		}
		defer fd.Close()

	}
}
