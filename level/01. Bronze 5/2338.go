package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	a, b := new(big.Int), new(big.Int)

	fmt.Fscanln(r, a)
	fmt.Fscanln(r, b)

	c := new(big.Int)
	fmt.Fprintln(w, c.Add(a, b))
	c = new(big.Int)
	fmt.Fprintln(w, c.Sub(a, b))
	c = new(big.Int)
	fmt.Fprintln(w, c.Mul(a, b))
}