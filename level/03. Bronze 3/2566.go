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
	var a, max, x, y int
	for i:=0;i<81;i++ {
		fmt.Fscan(r, &a)
		if a>max{
			max=a
			x= i/9+1
			y= i%9+1
		}
	}
	fmt.Fprintf(w,"%d\n%d %d",max,x,y)
}