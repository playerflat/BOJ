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
	var a int
	fmt.Fscanln(r, &a)
	if a == 1 {
		fmt.Fprintln(w, "*")
	} else {
		var xs, ys int
		xe := (a-1)*4 + 1
		ye := xe + 2

		z := make([][]string, ye)
		for i := 0; i < ye; i++ {
			in := make([]string, xe)
			for j := 0; j < xe; j++ {
				in[j] = " "
			}
			z[i] = in
		}

		for i := 0; i < a; i++ {
			draw(z, xs+i*2, ys+i*2, xe-i*2, ye-i*2)
		}
		z[ys+(a-1)*2+2][xe-(a-1)*2-2] = " "

		for j:=1;j<xe;j++{
			z[1][j]=""
		}

		for i := 0; i < ye; i++ {
			for j:=0;j<xe;j++{
				fmt.Fprint(w, z[i][j])
			}
			fmt.Fprintln(w, "")
		}
	}
}

func draw(z [][]string, xs, ys, xe, ye int) {

	for j := xe - 1; j >= xs; j-- {
		z[ys][j] = "*"
	}
	for i := ys; i < ye; i++ {
		z[i][xs] = "*"
	}
	for j := xs; j < xe; j++ {
		z[ye-1][j] = "*"
	}
	for i := ye - 1; i > ys+1; i-- {
		z[i][xe-1] = "*"
	}
	z[ys+2][xe-2] = "*"
}