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
	var a, b, c, d, e, f float64
	fmt.Fscanln(r, &a)
	for i := 0; i < int(a); i++ {
		fmt.Fscanf(r, "%f %f %f %f %f\n", &b, &c, &d, &e, &f)
		fmt.Fprintf(w, "$%.2f\n",b*350.34+c*230.90+d*190.55+e*125.30+f*180.90)
	}
}