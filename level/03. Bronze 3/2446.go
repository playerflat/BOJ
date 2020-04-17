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
	for i:=1;i<a;i++{
		RepeatPrint(i, a)
	}
	for i:=a;i>0;i--{
		RepeatPrint(i, a)
		if i==1{
			break
		}
	}
}

func RepeatPrint(i, a int){
	fmt.Fprint(w,strings.Repeat(" ",i-1))
	fmt.Fprint(w,strings.Repeat("*",2*a-(i*2-1)))
	fmt.Fprintln(w, "")
}