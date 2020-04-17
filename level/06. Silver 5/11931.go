package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a, b int
	c := make([]int, 0)
	fmt.Fscanln(r, &a)

	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		c = append(c, b)
	}
	
	sort.Ints(c)

	for i:=len(c)-1;i>=0;i--{
		fmt.Fprintln(w, c[i])
	}
}
