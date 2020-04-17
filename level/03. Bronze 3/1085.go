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
	var X, Y, W, H int
	fmt.Fscanf(r, "%d %d %d %d", &X, &Y, &W, &H)
	f := W - X
	b := H - Y

	if b < f {
		f = b
	}
	if X < f {
		f = X
	}
	if Y < f {
		f = Y
	}
	fmt.Fprintln(w, f)
}