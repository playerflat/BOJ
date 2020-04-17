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
	var a, b, c, d, e int
	fmt.Fscanf(r, "%d %d %d %d\n", &a, &b, &c, &d)
	for i:=1+b;i<=a;i+=b{
		e+=c
	}
	fmt.Fprintln(w, d*e)
}