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
	var n int
	fmt.Fscanln(r, &n)

	a := big.NewInt(1)
	b := big.NewInt(0)
	for i:=0;i<n;i++{
		c:= &big.Int{}
		a, b = b, c.Add(b, a)
	}

	fmt.Fprintln(w, b)
}