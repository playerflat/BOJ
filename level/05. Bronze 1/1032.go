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
	b := make([]byte, 0)
	var c []byte
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &b)
		if c == nil {
			c = b
		}
		for j := 0; j < len(b); j++ {
			if b[j] != c[j] {
				c[j] = '?'
			}
		}
	}
	fmt.Fprintln(w, string(c))
}
