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
	var a, b, c int
	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
	fmt.Fprintf(w, "%d %d",c/b, c%b)
}
