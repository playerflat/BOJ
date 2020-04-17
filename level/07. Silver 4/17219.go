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
	var id, pw string
	info := make(map[string]string)
	fmt.Fscanf(r, "%d %d\n", &a, &b)
	for i := 0; i < a; i++ {
		fmt.Fscanf(r, "%s %s\n", &id, &pw)
		info[id] = pw
	}
	for i := 0; i < b; i++ {
		fmt.Fscanln(r, &id)
		fmt.Fprintln(w, info[id])
	}
}
