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
	if a%2 == 0 {
		for i := 0; i < a; i++ {
			fmt.Fprintln(w, strings.Repeat("* ", a/2))
			fmt.Fprintln(w, strings.Repeat(" *", a/2))
		}
	} else {
		for i := 0; i < a; i++ {
			fmt.Fprintln(w, strings.Repeat("* ", a/2+1))
			fmt.Fprintln(w, strings.Repeat(" *", a/2))
		}
	}
}
