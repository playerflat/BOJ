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
	var R float64
	fmt.Fscanln(r, &R)
	fmt.Fprintf(w,"%.6f\n", math.Pi * math.Pow(R,2.0))
	fmt.Fprintf(w,"%.6f", 2*math.Pow(R,2.0))
}
