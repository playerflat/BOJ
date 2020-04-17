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
	var a, b string
	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	sum := big.NewInt(0)

	bi := &big.Int{}
	bi, _ = bi.SetString(a,10)
	sum = sum.Add(sum,bi)
	bi, _ = bi.SetString(b,10)
	sum = sum.Add(sum,bi)

	fmt.Fprintln(w, sum)
}