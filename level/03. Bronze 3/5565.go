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
	fmt.Fscanln(r, &a)
	for i := 0; i < 9; i++ {
		fmt.Fscanln(r, &b)
		a -= b
	}
	fmt.Fprintln(w, a)
}