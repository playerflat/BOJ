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
	var a, b, c, sum int
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		sum = 0
		fmt.Fscanln(r, &b)
		sum += b
		fmt.Fscanln(r, &c)
		for j := 0; j < c; j++ {
			var d, e int
			fmt.Fscanf(r, "%d %d\n", &d, &e)
			sum += d * e
		}
		fmt.Fprintln(w, sum)
	}
}
