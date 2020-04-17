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

	var a, max, c int
	for i := 1; i < 10; i++ {
		fmt.Fscanln(r, &a)
		if a>max{
			max=a
			c = i
		}
	}
	fmt.Fprintf(w,"%d\n%d",max,c)
}