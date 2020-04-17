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
	var n int
	fmt.Fscanln(r, &n)

	a, b := 1, 0
	for i:=0;i<n;i++{
		a, b = b, a+b
	}

	fmt.Fprintln(w, b)
}