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
	for i := 0; i < 5; i++ {
		fmt.Fscan(r, &a)
		b += math.Pow(a, 2)
	}
	fmt.Fprintln(w, int(b)%10)
}