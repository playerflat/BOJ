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
	fmt.Fscanln(r, &a)

	for i := 0; i < a; i++ {
		c := make([]int, 0)
		for j := 0; j < 10; j++ {
			fmt.Fscan(r, &b)
			c = append(c, b)
		}
		sort.Ints(c)
		fmt.Fprintln(w, c[7])
	}
}