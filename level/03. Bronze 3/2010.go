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

	for n!=0 {
		fmt.Fscanln(r, &a)
		b += a - 1
		n--
	}
	fmt.Fprintln(w, b+1)
}
