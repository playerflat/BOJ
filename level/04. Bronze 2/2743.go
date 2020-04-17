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
	var c string
	fmt.Fscanln(r, &c)
	fmt.Fprintln(w, len(c))
}