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
	var a string
	var b uint8
	var s int
	fmt.Fscanln(r, &a)
	for i:=0;i<len(a);i++{
		if b!= a[i]{
			s+=10
			b = a[i]
		} else{
			s+=5
		}
	}
	fmt.Fprintln(w, s)
}