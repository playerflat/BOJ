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
	for {
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		fmt.Fscan(r, &c)

		if a+b+c == 0 {
			break
		}

		a, b, c = a*a, b*b, c*c

		if a > b {
			a, b = b, a
		}

		if b > c {
			b, c = c, b
		}

		if a+b == c {
			fmt.Fprintln(w, "right")
		} else {
			fmt.Fprintln(w, "wrong")
		}

	}
}