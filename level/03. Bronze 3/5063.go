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
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d %d %d\n", &b, &c, &d)
		if b+d < c {
			fmt.Fprintln(w, "advertise")
		} else if b+d == c {
			fmt.Fprintln(w, "does not matter")
		} else {
			fmt.Fprintln(w, "do not advertise")
		}
	}
}