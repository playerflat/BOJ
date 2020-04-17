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
	var a string
	var sum int
	for i := 0; i < 8; i++ {
		fmt.Fscanln(r, &a)
		if i%2 == 0 {
			sum += check(a, 0) + check(a, 2) + check(a, 4) + check(a, 6)
		} else{
			sum += check(a, 1) + check(a, 3) + check(a, 5) + check(a, 7)
		}
	}
	fmt.Fprintln(w, sum)
}

func check(a string, n int) int {
	if a[n] == 70 {
		return 1
	}
	return 0
}