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
	fmt.Fscanln(r, &a)

	for i := a; i > 1; i-- {
		if i%5 == 0 {
			sum++
		}

		if i%25 == 0 {
			sum++
		}

		if i%125 == 0 {
			sum++
		}
	}
	fmt.Fprintln(w, sum)
}