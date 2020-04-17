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
	var a, sum int
	b := make([]int, 0)

	for i := 0; i < 5; i++ {
		fmt.Fscanln(r, &a)
		b = append(b, a)
		sum+=a
	}

	sort.Ints(b)

	fmt.Fprintln(w, sum/5)
	fmt.Fprintln(w, b[len(b)/2])
}