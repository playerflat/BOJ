package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	book := make(map[string]string)
	var a, b int
	var c string
	fmt.Fscanf(r, "%d %d\n", &a, &b)

	for i := 1; i <= a; i++ {
		fmt.Fscanln(r, &c)
		book[c] = strconv.Itoa(i)
		book[strconv.Itoa(i)] = c
	}

	for i := 0; i < b; i++ {
		fmt.Fscanln(r, &c)
		fmt.Fprintln(w, book[c])
	}
}