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

	var a, count int
	fmt.Fscanln(r, &a)

	for i := 0; i < a; i++ {
		count++
		z := make(map[string]int)
		var b string
		fmt.Fscanln(r, &b)
		for j:=0;j<len(b);j++{
			if _, t := z[string(b[j])]; t{
				if z[string(b[j])] != j-1{
					count--
					break
				} else{
					z[string(b[j])] = j
				}
			} else{
				z[string(b[j])] = j
			}
		}
	}
	fmt.Fprintln(w,count)
}