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
	for {
		b++
		c += b
		d++
		if c>a{
			d--
			break
		}
	}
	fmt.Fprintln(w, d)
}