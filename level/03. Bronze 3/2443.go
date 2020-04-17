package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a int
	fmt.Fscanln(r, &a)
	for i:=a;i>0;i-- {
		fmt.Fprint(w,strings.Repeat(" ",a-i))
		fmt.Fprint(w,strings.Repeat("*",i*2-1))
		if i==1{
			break
		}
		fmt.Fprintln(w,"")
	}
}