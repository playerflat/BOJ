package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var a, b, c, count int
	fmt.Fscanf(r, "%d %d %d", &a, &b, &c)

	if b >= c {
		count = -1
	} else {
		count=a/(c-b)+1
	}
	fmt.Fprintln(w, count)
}