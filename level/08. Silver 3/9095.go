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
	var n, x int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &x)
		a, b, c := 1, 0, 0
		for j := 0; j < x; j++ {
			a, b, c = a+b+c, a, b
		}
		fmt.Fprintln(w, a)
	}
}