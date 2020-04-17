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
	var n, a, b int
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		fmt.Fprintln(w, a*b/gcd(a,b))
	}
}

func gcd(x, y int) int {
	for y!=0{
		x, y = y, x%y
	}
	return x
}