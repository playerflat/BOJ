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

	var a, b, now, sum int
	var flag bool
	fmt.Fscanf(r, "%d %d\n", &a, &b)

	for i := 1; i <= b;{
		now += 1
		for j := 0; j < now; j++ {
			if i >= a && i <= b {
				sum += now
			}
			if i++; i > b {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Fprintln(w, sum)
}