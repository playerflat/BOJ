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
	fmt.Fscanf(r, "%d %d", &a, &b)
	x1, y1 := (a-1)/4, (a-1)%4
	x2, y2 := (b-1)/4, (b-1)%4

	fmt.Fprintln(w, Abs(x1-x2)+Abs(y1-y2))
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}