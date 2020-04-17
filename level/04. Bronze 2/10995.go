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
		fmt.Fprint(w, strings.Repeat("* ", a-1))
		fmt.Fprintln(w, "*")
		if i%2!=0{
			fmt.Fprint(w, " ")
		}
	}
}