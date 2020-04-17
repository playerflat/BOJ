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
	s := make([][]int, 0)
	s = append(s, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})
	for i:=0;i<14;i++{
		sn := make([]int, 0)
		var a int
		for j:=0; j<14;j++  {
			a += s[i][j]
			sn = append(sn, a)
		}
		s = append(s,sn)
	}
	
	var n int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(r, &a)
		fmt.Fscanln(r, &b)

		fmt.Fprintln(w, s[a][b-1])
	}
}