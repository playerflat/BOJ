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
	for i := 1; i <= a; i++ {
		if i != a {
			p(" ", a-i)
			p("*", 1)
			if j := (i-1)*2 - 1; j > 0 {
				p(" ", j)
			}
			if i != 1 {
				p("*", 1)
			}
			fmt.Fprintln(w, "")
		} else{
			p("*",a*2-1)
		}
	}
}

func p(a string, b int) {
	fmt.Fprint(w, strings.Repeat(a, b))
}
