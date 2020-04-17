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
	var a int
	fmt.Fscanln(r, &a)
	if a%2==1{
		fmt.Fprintln(w, "SK")
	} else{
		fmt.Fprintln(w, "CY")
	}
}