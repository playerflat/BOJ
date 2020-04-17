package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, b, l int
	for {
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		if a+b == 0 {
			break
		}
		if a > b {
			l = len(strconv.Itoa(a))
		} else {
			l = len(strconv.Itoa(b))
		}
		var c, d int
		for i := 0; i < l; i++ {
			if a%10+b%10+d > 9 {
				c++
				d = 1
			} else {
				d = 0
			}
			a /= 10
			b /= 10
		}
		fmt.Fprintln(w, c)
	}
}