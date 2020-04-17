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
	var a, b, c, count int
	fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
	for i := 1; ; i++ {
		count += a
		if i%7 == 0 {
			count += b
		}
		if count >= c {
			fmt.Fprintln(w, i)
			break
		}
	}
}