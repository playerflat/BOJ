package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()
	var a int
	var b string
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscanln(r, &b)
		c1,_:=strconv.Atoi(string(b[0]))
		c2,_:=strconv.Atoi(string(b[2]))
		fmt.Fprintln(w, c1+c2)
	}
}