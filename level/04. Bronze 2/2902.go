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
	fmt.Fscanln(r, &a)

	for i:=0;i<len(a);i++{
		if a[i]>=65 && a[i] <= 90{
			fmt.Fprint(w, string(a[i]))
		}
	}
}