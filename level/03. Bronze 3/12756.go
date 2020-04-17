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
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	fmt.Fscanf(r, "%d %d\n", &c, &d)
	for {
		b -= c
		d -= a

		if b<=0 && d<=0{
			fmt.Fprintln(w, "DRAW")
			break
		} else if b<=0{
			fmt.Fprintln(w, "PLAYER B")
			break
		} else if d<=0{
			fmt.Fprintln(w, "PLAYER A")
			break
		}
	}
}