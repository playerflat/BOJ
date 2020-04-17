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
	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	b+=a
	fmt.Fscanln(r, &a)
	a+=b
	fmt.Fscanln(r, &b)
	b+=a
	fmt.Fscanln(r, &a)
	a+=b
	fmt.Fprintln(w, a)
}