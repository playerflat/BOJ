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
	var a, b int
	min, max := 1000000, 1
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d ", &b)
		if b < min {
			min = b
		}
		if b > max {
			max = b
		}
	}
	fmt.Fprintln(w, min*max)
}