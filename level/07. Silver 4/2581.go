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
	var m, n, min, sum int
	fmt.Fscanln(r, &m)
	fmt.Fscanln(r, &n)
	for ; m <= n; m++ {
		if m ==1{
			continue
		}
		var b bool
		for a := 2; a < m; a++ {
			if m%a == 0 {
				b = true
				break
			}
		}
		if b {
			continue
		}
		sum += m
		if min == 0 {
			min = m
		}
	}

	if sum == 0 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, sum)
		fmt.Fprintln(w, min)
	}
}
