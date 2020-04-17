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
	fmt.Fscanln(r, &a)
	for i:=0;i<a;i++{
		var b, c int
		fmt.Fscanf(r,"%d %d\n",&b,&c)
		fmt.Fprintln(w,(b-c-2)*-1)
	}
}