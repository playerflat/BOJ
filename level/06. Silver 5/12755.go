package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a int
	var b string
	fmt.Fscanln(r, &a)
	i := 0
	for a > 0 {
		i++
		b = strconv.Itoa(i)
		a -= len(b)
		if a<=0{
			fmt.Fprintln(w, string(b[len(b)+a-1]))
		}
	}
}