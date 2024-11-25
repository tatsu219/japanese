package main

import (
	"fmt"
	"github.com/hasuburero/japanese/japanese"
	"io"
	"os"
	"strings"
)

const (
	input  = "dataset0.txt"
	output = "dataset1.txt"
)

const (
	o  = "お"
	u  = "う"
	du = "づ"
	zu = "ず"
	di = "ぢ"
	ji = "じ"
)

type output_struct struct {
	goji  string
	kanji string
}

func main() {
	fd_in, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Open error")
		return
	}
	defer fd_in.Close()
	input_buf, err := io.ReadAll(fd_in)
	if err != nil {
		fmt.Println(err)
		fmt.Println("io.ReadAll error")
		return
	}
	slice := strings.Split(string(input_buf), "\n")
	if slice[len(slice)-1] == "" {
		slice = slice[:len(slice)-1]
	}

	var output_buf []output_struct
	for _, ctx := range slice {
		original := strings.Split(ctx, ",")
		if len(original) < 2 {
			fmt.Println(err)
			fmt.Println("strings.Split error")
			return
		}
		buf := japanese.StrconvFirst(original[0], o, u)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
		buf = japanese.StrconvFirst(original[0], u, o)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
		buf = japanese.StrconvFirst(original[0], di, ji)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
		buf = japanese.StrconvFirst(original[0], ji, di)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
		buf = japanese.StrconvFirst(original[0], zu, du)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
		buf = japanese.StrconvFirst(original[0], du, zu)
		if buf != original[0] {
			output_buf = append(output_buf, output_struct{goji: buf, kanji: original[1]})
		}
	}

	fd_out, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		fmt.Println("os.Create error (output)")
		return
	}
	defer fd_out.Close()

	for i, ctx := range output_buf {
		fd_out.Write([]byte(ctx.goji))
		fd_out.Write([]byte(","))
		fd_out.Write([]byte(ctx.kanji))
		if i != len(output_buf)-1 {
			fd_out.Write([]byte("\n"))
		}
	}

	return
}
