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
	var a, b, c, d int
	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	fmt.Fscanln(r, &c)
	fmt.Fscanln(r, &d)

	fmt.Fprintln(w, (a+b+c+d)/60)
	fmt.Fprintln(w, (a+b+c+d)%60)
}