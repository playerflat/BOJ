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
	for i := 1; i < a; i++ {
		RepeatPrint(i)
	}
	for i := a;i>0;i--{
		RepeatPrint(i)
	}
}

func RepeatPrint(i int) {
	fmt.Fprintln(w, strings.Repeat("*", i))
}