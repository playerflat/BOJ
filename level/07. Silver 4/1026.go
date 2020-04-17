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
	mins, maxs := make([]int, 0), make([]int, 0)
	fmt.Fscanln(r, &a)

	mins = Append(a, mins)
	maxs = Append(a, maxs)

	sort.Ints(mins)
	sort.Sort(sort.Reverse(sort.IntSlice(maxs)))

	for i := 0; i < a; i++ {
		sum += mins[i] * maxs[i]
	}

	fmt.Fprintln(w, sum)

}

func Append(a int, s []int) []int {
	var b int
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d ", &b)
		s = append(s, b)
	}
	fmt.Fscanln(r, "")
	return s
}