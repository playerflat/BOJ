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
	var a string
	m := make(map[string]int)

	for i := 97; i <= 122; i++ {
		m[string(i)] = -1
	}

	fmt.Fscanln(r, &a)

	for i := len(a) - 1; i > -1; i-- {
		m[string(a[i])] = i
	}
	for i := 97; i <= 122; i++ {
		fmt.Fprintf(w, "%d", m[string(i)])
		if i != 122 {
			fmt.Fprint(w," ")
		}
	}
}