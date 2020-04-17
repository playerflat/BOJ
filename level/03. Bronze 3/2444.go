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
		fmt.Fprint(w, strings.Repeat(" ", a-i))
		fmt.Fprint(w, strings.Repeat("*", i+i-1))
		fmt.Fprintln(w, "")
	}

	for i := 1; i < a; i++ {
		fmt.Fprint(w, strings.Repeat(" ", i))
		fmt.Fprint(w, strings.Repeat("*", a-i+a-i-1))
		fmt.Fprintln(w, "")
	}
}