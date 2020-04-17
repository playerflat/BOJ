package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var n int
	fmt.Fscanln(r,&n)
	for i:=0; i<n; i++{
		var a,b,c,j int
		fmt.Fscan(r, &a)
		fmt.Fscan(r, &b)
		fmt.Fscan(r, &c)

		if c%a == 0{
			j=a
		} else{
			j = c%a
		}
		fmt.Fprint(w, j)

		fmt.Fprintf(w, "%02d\n", int(math.Ceil(float64(c)/float64(a))))
	}
}