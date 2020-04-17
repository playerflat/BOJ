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

	var a int
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		var b int
		var c string
		fmt.Fscanln(r, &b, &c)
		for j := 0; j < len(c); j++ {
			fmt.Fprintf(w, "%s", strings.Repeat(string(c[j]), b))
		}
		fmt.Fprintln(w,"")
	}
}