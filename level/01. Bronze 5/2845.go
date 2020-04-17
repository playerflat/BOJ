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
	var a, b, c, d, e, f, g int
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	fmt.Fscanf(r, "%d %d %d %d %d\n", &c, &d, &e, &f, &g)
	a *= b
	fmt.Fprintf(w, "%d %d %d %d %d", c-a, d-a, e-a, f-a, g-a)
}