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
	var a, b, x, y int
	fmt.Fscanf(r, "%d %d", &x, &y)

	a = x
	b = y

	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}

	fmt.Fprintln(w,a)

	a = x/a
	b = y
	fmt.Fprintln(w, b*a)
}