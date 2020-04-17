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
	var a, b, c, count int

	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)

	for {
		if b-a > (b-c)*-1 {
			count++
			b, c = a+1, b
		} else if b-a < (b-c)*-1 {
			count++
			a, b = b, b+1
		} else if b-a == (b-c)*-1 {
			if a+1 == b && b == c-1 {
				break
			} else {
				count++
				a, b = b, b+1
			}
		}
	}
	fmt.Fprintln(w, count)
}