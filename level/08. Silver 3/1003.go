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
	var a, b, c int
	fib := make([]int, 0)
	a = 1
	for i := 0; i <= 40; i++ {
		b = a
		a = c
		c = a + b
		fib = append(fib, c)
	}
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		if b == 0 {
			fmt.Fprintln(w, "1 0")
		} else if b == 1 {
			fmt.Fprintln(w, "0 1")
		} else {
			fmt.Fprintln(w, fib[b-2], fib[b-1])
		}
	}
}