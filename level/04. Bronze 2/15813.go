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
	var a, sum int
	var b string
	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	for i:=0; i<a;i++  {
		sum += int(b[i]-64)
	}
	fmt.Fprintln(w, sum)
}