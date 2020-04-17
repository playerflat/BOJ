package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

type loc struct {
	x int
	y int
}

type sloc []loc

func (s sloc) Len() int {
	return len(s)
}

func (s sloc) Less(i, j int) bool {
	if s[i].x == s[j].x {
		return s[i].y < s[j].y
	} else {
		return s[i].x < s[j].x
	}
}

func (s sloc) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	defer w.Flush()
	var a, x, y int
	b := make([]loc, 0)
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d %d\n", &x, &y)
		b = append(b, loc{x, y})
	}

	sort.Sort(sloc(b))

	for i:=0;i<len(b);i++{
		fmt.Fprintf(w,"%d %d\n", b[i].x,b[i].y)
	}
}