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
	var a, b int
	fmt.Fscanln(r, &a)

	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		var sum int
		z := make([]int, 10)
		l := len(strconv.Itoa(b))
		for j := 0; j < l; j++ {
			c := b % 10
			b /= 10
			z[c]++
		}

		for _, v := range z {
			if v > 0 {
				sum++
			}
		}

		fmt.Fprintln(w, sum)
	}
}