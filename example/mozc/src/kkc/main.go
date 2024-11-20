package main

import (
	"errors"
	"fmt"
	"github.com/hasuburero/japanese/japanese"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	mozc_dictionary_dir         = "../../mozc-dictionary/"
	mozc_dictionary_name        = "dictionary"
	mozc_dictionary_count       = 10
	additional_dictionary_dir   = "../../additional-dictionary/"
	additional_dictionary_name  = "additinal2"
	additional_dictionary_count = 12
)

const (
	port = ":18080"
)

type mozc_struct struct {
	yomi             string
	left_context_id  int16
	right_context_id int16
	cost             int16
	kanji            string
}

var dictionary map[string][]mozc_struct

func addDictionary(fd *os.File) error {
	buf, err := io.ReadAll(fd)
	if err != nil {
		return err
	}
	line_buf := strings.Split(string(buf), "\n")
	if line_buf[len(line_buf)-1] == "" {
		if len(line_buf) != 1 {
			line_buf = line_buf[:len(line_buf)-1]
		}
	}
	for _, ctx := range line_buf {
		slice := strings.Split(ctx, "\t")
		yomi := slice[0]
		left_context_id, err := strconv.Atoi(slice[1])
		right_context_id, err := strconv.Atoi(slice[2])
		cost, err := strconv.Atoi(slice[3])
		kanji := slice[4]
		if err != nil {
			return err
		}
		dictionary[yomi] = append(dictionary[yomi], mozc_struct{yomi: yomi, left_context_id: int16(left_context_id), right_context_id: int16(right_context_id), cost: int16(cost), kanji: kanji})
	}

	return nil
}

func readDictionary() error {
	for i := range mozc_dictionary_count {
		filename := mozc_dictionary_dir + mozc_dictionary_name + fmt.Sprintf("%02d", i) + ".txt"
		fd, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer fd.Close()

		addDictionary(fd)

	}

	return nil
}

func main() {
	fmt.Println("KKC Server")
}
