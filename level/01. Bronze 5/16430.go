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
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	fmt.Fprintf(w, "%d %d", b-a, b)
}
