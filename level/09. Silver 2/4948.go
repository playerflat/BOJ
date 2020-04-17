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
	for {
		var count int
		fmt.Fscanln(r, &a)
		if a == 0 {
			break
		} // 3 받으면 4 ~ 6 사이의 소수 개수 구하
		b := a * 2
		a++
		for ; a <= b; a++ {
			if a != 2 && a%2 == 0 {
				continue
			}
			var b bool
			for x := 3; x*x <= a; x++ {
				if a%x == 0 {
					b = true
					break
				}
			}
			if b {
				continue
			}
			count++
		}
		fmt.Fprintln(w, count)
	}
}