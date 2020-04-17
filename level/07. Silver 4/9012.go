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
	var n int
	var x string
	fmt.Fscanln(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &x)
		var count int
		for j := 0; j < len(x); j++ {
			if string(x[j]) == "(" {
				count++
			} else if string(x[j]) == ")" {
				count--
			}

			if count<0{
				break
			}
		}
		if count==0{
			fmt.Fprintln(w,"YES")
		}else{
			fmt.Fprintln(w, "NO")
		}
	}
}