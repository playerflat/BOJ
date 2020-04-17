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
	var a, b, sum int
	fmt.Fscanln(r, &a)
	for i := 0; i < 5; i++ {
		fmt.Fscan(r, &b)
		if a == b {
			sum += 1
		}
	}
	fmt.Fprintln(w, sum)
}