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
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &b)
		fmt.Fscan(r, &c)
		fmt.Fprintln(w, b*c/gcd(b, c), gcd(b, c))
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}