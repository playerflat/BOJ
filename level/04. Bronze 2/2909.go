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
	var a, b float64
	fmt.Fscanf(r, "%f %f\n", &a, &b)
	fmt.Fprintf(w,"%.f", math.Round(a/math.Pow(10,b))*math.Pow(10,b))
}