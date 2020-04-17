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
	var x int
	fmt.Fscanf(r, "%x", &x)
	fmt.Fprintln(w, x)
}