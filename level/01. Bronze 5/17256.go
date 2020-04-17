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
	a := make([]int, 3)
	c := make([]int, 3)

	fmt.Fscanf(r, "%d %d %d\n", &a[0], &a[1], &a[2])
	fmt.Fscanf(r, "%d %d %d\n", &c[0], &c[1], &c[2])

	fmt.Fprintf(w, "%d %d %d",c[0]-a[2],c[1]/a[1],c[2]-a[0])
}