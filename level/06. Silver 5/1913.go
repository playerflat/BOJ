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
	var a, b, xs, ys, xe, ye, gx,gy int
	var flag bool
	fmt.Fscanln(r, &a)
	fmt.Fscanln(r, &b)
	snail := make([][]int, a)
	for i:=0;i<a;i++{
		for j:=0;j<a;j++{
			snail[i] = append(snail[i], 0)
		}
	}
	xe, ye = a, a
	c := a * a
	for {
		for i := xs; i < xe; i++ {
			snail[i][ys] = c
			if c==b{
				gx, gy = i, ys
			}
			c--
			if c == 0 {
				flag = true
				break
			}
		}
		ys++
		xe--

		if flag {
			break
		}

		for j := ys; j < ye; j++ {
			snail[xe][j] = c
			if c==b{
				gx, gy = xe, j
			}
			c--
			if c == 0 {
				flag = true
				break
			}
		}
		ye--

		if flag {
			break
		}

		for i := xe - 1; i >= xs; i-- {
			snail[i][ye] = c
			if c==b{
				gx, gy = i, ye
			}
			c--
			if c == 0 {
				flag = true
				break
			}
		}

		if flag {
			break
		}

		for j := ye - 1; j >= ys; j-- {
			snail[xs][j] = c
			if c==b{
				gx, gy = xs, j
			}
			c--
			if c == 0 {
				flag = true
				break
			}
		}
		xs++

		if flag {
			break
		}

	}

	for i:=0;i<a;i++{
		for j:=0;j<a;j++{
			fmt.Fprintf(w,"%d ", snail[i][j])
		}
		fmt.Fprintln(w, "")
	}
	fmt.Fprintln(w, gx+1, gy+1)
}