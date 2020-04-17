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
	var n, count int
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		var a int
		var b bool
		fmt.Fscan(r, &a)
		if a==1{
			continue
		}
		for j := 2; j < a; j++ {
			if a%j == 0 {
				b=true
				break
			}
		}
		if b{
			continue
		}
		count++
	}
	fmt.Fprintln(w, count)
}
