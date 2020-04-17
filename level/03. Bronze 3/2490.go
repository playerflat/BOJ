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
	for i := 0; i < 3; i++ {
		fmt.Fscanf(r, "%d %d %d %d\n", &a, &b, &c, &d)
		switch a + b + c + d {
		case 0:
			fmt.Fprintln(w, "D")
		case 1:
			fmt.Fprintln(w, "C")
		case 2:
			fmt.Fprintln(w, "B")
		case 3:
			fmt.Fprintln(w, "A")
		case 4:
			fmt.Fprintln(w, "E")
		}
	}
}