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
	var m, n int
	fmt.Fscanf(r, "%d %d", &m, &n)
	for ; m <= n; m++ {
		if m == 1{
			continue
		}

		if m!=2 && m%2==0{
			continue
		}

		var b bool

		for a := 3; a*a<=m; a++ {
			if m%a == 0 {
				b = true
				break
			}
		}
		if b {
			continue
		}
		fmt.Fprintln(w, m)
	}
}