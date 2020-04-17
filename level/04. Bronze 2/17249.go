package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a string
	fmt.Fscanln(r, &a)
	b :=strings.Split(a,"0")
	fmt.Fprintln(w,strings.Count(b[0],"@"), strings.Count(b[1],"@"))
}