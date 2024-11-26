package main

import (
	"fmt"
	"github.com/hasuburero/japanese/japanese"
	"io"
	"os"
	"strconv"
	"strings"
	//"errors"
	//"net/http"
)

const (
	dataset_dir   = "../dataset/"
	dataset_name  = "dataset"
	dataset_count = 2
)

const (
	mozc_dictionary_dir         = "../../../mozc-dictionary/"
	mozc_dictionary_name        = "dictionary"
	mozc_dictionary_count       = 10
	additional_dictionary_dir   = "../../../additional-dictionary/"
	additional_dictionary_name  = "additional2"
	additional_dictionary_count = 12
	connection_def              = "connection.txt"
)

const (
	port = ":18080"
)

type mozc_struct struct {
	yomi             string
	left_context_id  int
	right_context_id int
	cost             int
	kanji            string
	rune_length      int
}

type node_struct struct {
	mozc  mozc_struct
	start int
	end   int
}

type node_info struct {
	cost   int
	parent int
}

type test_struct struct {
	kana  string
	kanji string
}

var dictionary map[string][]mozc_struct
var connection_cost [][]int

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
		dictionary[yomi] = append(dictionary[yomi], mozc_struct{yomi: yomi, left_context_id: int(left_context_id), right_context_id: int(right_context_id), cost: int(cost), kanji: kanji, rune_length: japanese.RuneLength(yomi)})
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
	if slice[len(slice)-1] == "" {
		slice = slice[:len(slice)-1]
	}
	width, err := strconv.Atoi(slice[0])
	if err != nil {
		return err
	}
	slice = slice[1:]

	height := 0
	connection_cost = append(connection_cost, []int{})
	for i, ctx := range slice {
		buf, err := strconv.Atoi(ctx)
		if err != nil {
			return err
		}
		connection_cost[height] = append(connection_cost[height], int(buf))

		if i%width == 0 && i != len(slice)-1 && i != 0 {
			connection_cost = append(connection_cost, []int{})
			height++
		}
	}

	return nil
}

func searchTango(arg string) ([]mozc_struct, bool, error) {
	ctx, exists := dictionary[arg]
	if !exists {
		return []mozc_struct{}, exists, nil
	}
	return ctx, exists, nil
}

func reverseArray(arg []int, index int) []int {
	if index == len(arg)-1 {
		return []int{index}
	} else {
		return append(reverseArray(arg, index+1), arg[index])
	}
}

func dijkstra(node_list [][]int) []int {
	node_size := len(node_list)
	var node_buf []int
	var node []node_info = make([]node_info, node_size)

	for i, ctx := range node_list[0] {
		if ctx != 0 {
			node[i].cost = node[0].cost + node_list[0][i]
			node[i].parent = 0
			node_buf = append(node_buf, i)
		}
	}
	for len(node_buf) != 0 {
		target := node_buf[0]
		node_buf = node_buf[1:]
		for i, ctx := range node_list[target] {
			if ctx == -1 {
				cost_buf := node[target].cost
				if node[i].cost == 0 || node[i].cost > cost_buf {
					node[i].cost = cost_buf
					node[i].parent = target
				}
			} else if ctx != 0 {
				cost_buf := node[target].cost + ctx
				if node[i].cost == 0 || node[i].cost > cost_buf {
					node[i].cost = cost_buf
					node[i].parent = target
					node_buf = append(node_buf, i)
				}
			}
		}
	}

	var index int = node[len(node)-1].parent
	var result []int = []int{index}
	for index != 0 {
		index = node[index].parent
		result = append(result, index)
	}
	result = reverseArray(result, 0)

	return result
}

func converter(arg string) (string, error) {
	arg_length := japanese.RuneLength(arg)
	tango_array := []node_struct{}
	for i := range arg_length {
		for j := i + 1; j <= arg_length; j++ {
			tango, err := japanese.RuneSubstring(arg, i, j)
			if err != nil {
				return "", err
			}
			mozc_array, exists, err := searchTango(tango)
			if err != nil {
				return "", err
			} else if !exists {
				continue
			}
			for _, ctx := range mozc_array {
				tango_array = append(tango_array, node_struct{ctx, i, j})
			}
		}
	}

	tango_array_length := len(tango_array)
	node_list := make([][]int, tango_array_length+2)
	for i := range tango_array_length + 2 {
		node_list[i] = make([]int, tango_array_length+2)
	}
	for i, ctx := range tango_array {
		if ctx.start == 0 {
			cost := connection_cost[0][ctx.mozc.left_context_id] + ctx.mozc.cost
			node_list[0][i+1] = cost
		}
	}

	for i, ctx := range tango_array {
		if ctx.end == arg_length {
			node_list[i+1][len(tango_array)+1] = -1
			continue
		}
		for j := 0; j < len(tango_array); j++ {
			if tango_array[j].start == ctx.end {
				cost := connection_cost[ctx.mozc.right_context_id][tango_array[j].mozc.left_context_id] + tango_array[j].mozc.cost
				node_list[i+1][j+1] = cost
			}
		}
	}

	node_buf := dijkstra(node_list)
	var result string = ""
	//var info
	for i := 1; i < len(node_buf); i++ {
		result += tango_array[node_buf[i]-1].mozc.kanji
	}

	return result, nil
}

func main() {
	fmt.Println("KKC Server")
	dictionary = make(map[string][]mozc_struct)
	err := readDictionary()
	if err != nil {
		fmt.Println(err)
		fmt.Println("readDictionary error")
		return
	}
	fmt.Println("end of reading dictionary")

	err = readConnection()
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Open error connection_def file")
		return
	}
	fmt.Println("end of connection file")

	var dataset_num float32
	var dataset []test_struct
	for i := range dataset_count {
		filename := dataset_dir + dataset_name + strconv.Itoa(i) + ".txt"
		fd, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("os.Open error (dataset)")
		}
		defer fd.Close()

		content, err := io.ReadAll(fd)
		if err != nil {
			fmt.Println(err)
			fmt.Println("io.ReadAll error (dataset)", i)
			return
		}
		slice := strings.Split(string(content), "\n")
		if slice[len(slice)-1] == "" {
			slice = slice[:len(slice)-1]
		}
		for _, ctx := range slice {
			dataset_num += 1.0
			buf := strings.Split(ctx, ",")
			dataset = append(dataset, test_struct{buf[0], buf[1]})
		}
	}

	var correct_count float32 = 0.0
	for _, ctx := range dataset {
		result, err := converter(ctx.kana)
		if err != nil {
			fmt.Println(err)
			fmt.Println("converter error")
			return
		}
		fmt.Printf("%s, %s, %s\n", ctx.kana, result, ctx.kanji)
		if result == ctx.kanji {
			correct_count += 1.0
		}
	}
	fmt.Printf("%f / %f = %f\n", correct_count, dataset_num, correct_count/dataset_num)

	return
}
