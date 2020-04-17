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
	var a, b, c, d, e int

	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	fmt.Fscanln(r, &c)
	fmt.Fscanln(r, &d)
	fmt.Fscanln(r, &e)

	if a < b {
		b = a
	}

	if b > c {
		b = c
	}

	if d > e {
		d = e
	}

	fmt.Fprintln(w, b+d-50)
}
