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

	var x, y int
	fmt.Fscanf(r, "%d %d", &x, &y)
	fmt.Fprintln(w, Rev(Rev(x)+Rev(y)))
}

func Rev(n int) int {
	var t int
	for {
		t += n % 10
		n = n / 10
		if n==0{
			break
		}
		t *= 10
	}
	return t
}