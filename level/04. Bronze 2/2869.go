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
	var a, b, v int
	fmt.Fscanf(r, "%d %d %d", &a, &b, &v)
	fmt.Fprintln(w, (v-b-1)/(a-b)+1)
}