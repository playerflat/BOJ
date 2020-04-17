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
	var n int
	fmt.Fscanln(r,&n)
	fmt.Fprintln(w,fact(n))

}

func fact(n int) int{
	if n>0{
		return n*fact(n-1)
	}
	return 1
}