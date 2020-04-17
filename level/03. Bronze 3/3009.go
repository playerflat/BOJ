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
	x := make(map[int]int)
	y := make(map[int]int)
	var a,b int
	for i:=0; i<3; i++{
		fmt.Fscan(r,&a)
		fmt.Fscan(r,&b)

		if _, exist := x[a]; exist{
			delete(x,a)
		} else if !exist{
			x[a] = 1
		}

		if _, exist := y[b]; exist{
			delete(y,b)
		} else if !exist{
			y[b] = 1
		}
	}

	for a := range(x){
		fmt.Fprintf(w,"%d ",a)
	}
	for b:= range(y){
		fmt.Fprint(w,b)
	}
}