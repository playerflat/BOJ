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
	b := make(map[string]int)

	for i:=97;i<=122;i++{
		b[string(i)]=0
	}

	fmt.Fscanln(r, &a)

	for i:=0;i<len(a);i++{
		b[string(a[i])] += 1
	}

	for i:=97;i<=122;i++{
		fmt.Fprintf(w,"%d ",b[string(i)])
	}
}