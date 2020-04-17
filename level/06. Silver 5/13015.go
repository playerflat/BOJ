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
		p(i, a)
	}

	for i := a - 1; i > 0; i-- {
		p(i, a)
	}
}

func p(i, a int) {
	if i == 1 {
		fmt.Fprint(w, strings.Repeat("*", a))
		fmt.Fprint(w, strings.Repeat(" ", (a*2)-1-i*2))
		fmt.Fprintln(w, strings.Repeat("*", a))
	} else if i == a {
		fmt.Fprint(w, strings.Repeat(" ", i-1))
		fmt.Fprint(w, "*")
		fmt.Fprint(w, strings.Repeat(" ", a-2))
		fmt.Fprint(w, "*")
		fmt.Fprint(w, strings.Repeat(" ", a-2))
		fmt.Fprintln(w, "*")
	} else {
		fmt.Fprint(w, strings.Repeat(" ", i-1))
		fmt.Fprint(w, "*")
		fmt.Fprint(w, strings.Repeat(" ", a-2))
		fmt.Fprint(w, "*")
		fmt.Fprint(w, strings.Repeat(" ", (a*2)-1-i*2))
		fmt.Fprint(w, "*")
		fmt.Fprint(w, strings.Repeat(" ", a-2))
		fmt.Fprintln(w, "*")
	}
}
