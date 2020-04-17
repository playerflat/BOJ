package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a float64
	var b, sum int
	c := make([]int,0)
	fmt.Fscanln(r, &a)
	for i := 0; i < int(a); i++ {
		fmt.Fscanln(r, &b)
		c = append(c, b)
	}

	d := math.Round(a/100*15)

	sort.Ints(c)
	for i:=d; i<a-d;i++{
		sum+= c[int(i)]
	}

	if a==0{
		fmt.Fprintln(w, 0)
	} else{
		fmt.Fprintln(w, math.Round(float64(sum)/(a-2*d)))
	}
}