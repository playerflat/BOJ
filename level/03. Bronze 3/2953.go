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
	for i := 1; i <= 5; i++ {
		b = 0
		for j := 0; j < 4; j++ {
			fmt.Fscan(r, &a)
			b += a
		}
		fmt.Fscanln(r, &a)
		if b > c {
			c = b
			d = i
		}
	}
	fmt.Fprintln(w, d, c)
}
