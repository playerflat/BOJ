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
	var a, b int
	fmt.Fscanln(r, &a)

	for i:=0;i<a;i++{
		fmt.Fscanln(r, &b)
		for j:=0;j<b;j++{
			for k:=0;k<b;k++{
				if j==0 || j==b-1 || k==0 || k==b-1 {
					fmt.Fprint(w, "#")
				} else{
					fmt.Fprint(w, "J")
				}
			}
			fmt.Fprintln(w, "")
		}
		fmt.Fprintln(w, "")
	}
}