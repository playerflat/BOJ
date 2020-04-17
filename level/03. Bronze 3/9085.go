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
		fmt.Fscanln(r, &b)
		for j := 0; j < b; j++ {
			fmt.Fscan(r, &c)
			sum += c
		}
		fmt.Fscanln(r, &c)
		fmt.Fprintln(w, sum)
		sum = 0
	}
}