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
	var a, b, c, d, e, f int
	fmt.Fscanf(r, "%d %d %d %d %d %d\n", &a, &b, &c, &d, &e, &f)
	fmt.Fprintf(w, "%d %d %d %d %d %d", 1-a, 1-b, 2-c, 2-d, 2-e, 8-f)
}