package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	//"github.com/hasuburero/japanese/japanese"
	//"net/http"
)

const (
	mozc_dictionary_dir         = "../../mozc-dictionary/"
	mozc_dictionary_name        = "dictionary"
	mozc_dictionary_count       = 10
	additional_dictionary_dir   = "../../additional-dictionary/"
	additional_dictionary_name  = "additional2"
	additional_dictionary_count = 12
	connection_def              = "connection.txt"
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
var connection_cost [][]int16

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
		if err != nil {
			return err
		}
	}

	for i := range additional_dictionary_count {
		filename := additional_dictionary_dir + additional_dictionary_name + fmt.Sprintf("%02d", i) + ".txt"
		fd, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer fd.Close()

		err = addDictionary(fd)
		if err != nil {
			return err
		}
	}

	return nil
}

func readConnection() error {
	fd, err := os.Open(mozc_dictionary_dir + connection_def)
	if err != nil {
		return err
	}
	defer fd.Close()

	content, err := io.ReadAll(fd)
	if err != nil {
		return err
	}

	slice := strings.Split(string(content), "\n")
	conn_width, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	slice = slice[1:]

	height := 0
	for i, ctx := range slice {
		i % conn_width
	}

	return nil
}

func converter() {

}

func main() {
	fmt.Println("KKC Server")
	stdin := bufio.NewScanner(os.Stdin)
	dictionary = make(map[string][]mozc_struct)
	err := readDictionary()
	if err != nil {
		fmt.Println(err)
		fmt.Println("readDictionary error")
		return
	}
	for _, dic_ctx := range dictionary {
		fmt.Println(dic_ctx[0])
		break
	}

	err = readConnection()
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Open error connection_def file")
		return
	}
	defer fd.Close()

	for {
		stdin.Scan()
		fmt.Println(stdin.Text())

		fmt.Println(dictionary[stdin.Text()])

		for _, ctx := range stdin.Text() {
			fmt.Print(ctx)
			fmt.Printf(" %x\n", ctx)
		}
	}
}
