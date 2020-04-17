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
	var a string
	fmt.Fscanln(r, &a)
	b := new(big.Int)
	b, _ = b.SetString(a, 2)
	fmt.Fprintf(w, "%o", b)
}