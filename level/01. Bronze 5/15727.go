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
	var a int
	fmt.Fscanln(r, &a)
	if a%5 == 0 {
		fmt.Fprintln(w, a/5)
	} else {
		fmt.Fprintln(w, a/5+1)
	}
}