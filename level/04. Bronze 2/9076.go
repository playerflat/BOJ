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
	var a int
	z := make([]int, 5)
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%d %d %d %d %d\n", &z[0], &z[1], &z[2], &z[3], &z[4])
		sort.Ints(z)
		if z[3]-z[1]>=4{
			fmt.Fprintln(w, "KIN")
		} else{
			fmt.Fprintln(w,z[1]+z[2]+z[3])
		}
	}
}