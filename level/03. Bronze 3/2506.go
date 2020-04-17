package main

import (
	"bufio"
	"fmt"
	"os"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main(){
	defer w.Flush()

	var a, b, c, sum int
	fmt.Fscanln(r, &a)

	for i:=0; i<a;  i++{
		fmt.Fscanf(r, "%d ", &b)
		if b==1{
			c++
			sum+=c
		} else{
			c=0
		}
	}
	fmt.Fprintln(w, sum)
}
