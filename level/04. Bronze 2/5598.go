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
	// A65 Z90
	var a string
	var b uint8
	fmt.Fscanln(r, &a)

	for i := 0; i < len(a); i++ {
		b = a[i]-3
		if b<65{
			b+=26
		}

		fmt.Fprint(w, string(b))
	}
}