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
		fmt.Fprintln(w, "I LOVE CBNU")
	} else {
		fmt.Fprintln(w, strings.Repeat("*", a))
		fmt.Fprint(w, strings.Repeat(" ", a/2))
		fmt.Fprintln(w, "*")
		for i := 0; i < a/2; i++ {
			fmt.Fprint(w, strings.Repeat(" ", a/2-i-1))
			fmt.Fprint(w, "*")
			fmt.Fprint(w, strings.Repeat(" ",i*2+1))
			fmt.Fprintln(w, "*")
		}
	}
}