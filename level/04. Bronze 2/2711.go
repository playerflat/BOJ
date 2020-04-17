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
	var a, b int
	var c string
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &b)
		fmt.Fscan(r, &c)
		fmt.Fprintln(w, strings.Join(strings.Split(c,"")[:b-1],"")+strings.Join(strings.Split(c,"")[b:],""))
	}
}