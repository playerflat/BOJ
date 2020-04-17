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
	count, m, min, max := 1, 0, 2, 7

	fmt.Fscanln(r, &a)
	for{
		if a==1{
			break
		} else{
			count +=1
			if min+m*6 <= a && a <= max+m*6{
				break
			}
			min += m*6
			m +=1
			max += m*6
		}
	}
	fmt.Fprintln(w, count)
}