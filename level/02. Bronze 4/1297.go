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
	var a, b, c float64
	fmt.Fscanf(r, "%f %f %f\n", &a, &b, &c)
	a = math.Pow(a, 2)
	d := math.Pow(a/(math.Pow(b, 2)+math.Pow(c, 2)), 0.5)

	fmt.Fprintf(w, "%.f %.f", math.Floor(b*d), math.Floor(c*d))
}