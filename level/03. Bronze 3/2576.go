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
	var a, sum int
	min := 100
	for i:=0;i<7;i++ {
		fmt.Fscan(r, &a)
		if a%2 != 0{
			if min>a{
				min=a
			}
			sum+=a
		}
	}
	if sum !=0{
		fmt.Fprintln(w, sum)
		fmt.Fprintln(w, min)
	} else{
		fmt.Fprintln(w,-1)
	}
}