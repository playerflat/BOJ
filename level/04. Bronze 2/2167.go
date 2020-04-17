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
	z := make([][]int, 0)
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	for i := 0; i < a; i++ {
		inner := make([]int, 0)
		for j := 0; j < b; j++ {
			fmt.Fscanf(r, "%d ", &c)
			inner = append(inner, c)
		}
		fmt.Fscanln(r, "")
		z = append(z, inner)
	}

	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		var sum int
		fmt.Fscanf(r, "%d %d %d %d\n", &b, &c, &d, &e)
		for j := b; j <= d; j++ {
			for k := c; k <= e; k++ {
				sum += z[j-1][k-1]
			}
		}
		fmt.Fprintln(w, sum)
	}
}