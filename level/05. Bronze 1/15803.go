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
	var x1,y1,x2,y2,x3,y3 int

	fmt.Fscanln(r, &x1, &y1)
	fmt.Fscanln(r, &x2, &y2)
	fmt.Fscanln(r, &x3, &y3)

	res := x1*y2+x2*y3+x3*y1;
	res = res - y1*x2-y2*x3-y3*x1

	if res == 0 {
		fmt.Fprintln(w, "WHERE IS MY CHICKEN?")
	} else {
		fmt.Fprintln(w, "WINNER WINNER CHICKEN DINNER!")
	}
}