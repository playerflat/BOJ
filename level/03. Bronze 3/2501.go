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
	c := make([]int, 0)
	fmt.Fscanf(r, "%d %d\n", &a, &b)

	for i := 1; i <= a/2; i++ {
		if a%i == 0 {
			c = append(c, i)
		}
	}
	c = append(c, a)

	if len(c)<b{
		fmt.Fprintln(w, 0)
	} else{
		fmt.Fprintln(w, c[b-1])
	}
}
