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
	var a, c, sum int
	b := make([]int, 100)

	for i := 0; i < 10; i++ {
		fmt.Fscanln(r, &a)
		b[a/10]++
		if b[c] < b[a/10] {
			c = a/10
		}
		sum+=a
	}

	fmt.Fprintln(w, sum/10)
	fmt.Fprintln(w, c*10)
}