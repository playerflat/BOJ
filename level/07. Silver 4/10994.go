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
	var a, b int
	fmt.Fscanln(r, &a)
	length := a*4 - 3
	z := make([][]string, length)
	for i := 0; i < length; i++ {
		in := make([]string, length)
		z[i] = in
	}

	for i := a; i > 0; i-- {
		draw(z, length, i, b)
		b += 2
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Fprint(w, z[i][j])
		}
		fmt.Fprintln(w, "")
	}
}

func draw(z [][]string, length, a, b int) {
	if a >= 1 {
		end := length - b
		for i := b; i < end; i++ {
			for j := b; j < end; j++ {
				z[i][j] = "*"
			}
		}
		for i := b + 1; i < end-1; i++ {
			for j := b + 1; j < end-1; j++ {
				z[i][j] = " "
			}
		}
	} else {
		z[b][b] = "*"
	}
}