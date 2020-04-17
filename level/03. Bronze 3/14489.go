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
	var a, b, c int
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	fmt.Fscanln(r, &c)

	{
		d := a + b
		e := c * 2
		if d >= e {
			fmt.Fprintln(w, d-e)
		} else{
			fmt.Fprintln(w, d)
		}
	}
}