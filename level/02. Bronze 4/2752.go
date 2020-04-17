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
	var a, b, c int
	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
	d := []int{a, b, c}

	sort.Ints(d)
	fmt.Fprintf(w, "%d %d %d\n", d[0], d[1], d[2])
}