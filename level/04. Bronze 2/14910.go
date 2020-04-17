package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a int
	b := make([]int, 0)
	for {
		_, err := fmt.Fscan(r, &a)
		if err == io.EOF {
			break
		} else {
			b = append(b, a)
		}
	}

	prev := -1000001
	flag := true
	for i := 0; i < len(b); i++ {
		if prev > b[i] {
			flag = false
			break
		}
		prev = b[i]
	}

	if flag {
		fmt.Fprintln(w, "Good")
	} else {
		fmt.Fprintln(w, "Bad")
	}
}