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
	var n int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		var a, b float64
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		c := 1.0

		for ; a > 0; a-- {
			c *= b / a
			b--
		}
		fmt.Fprintf(w, "%.0f\n", c)
	}
}