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

	var n int
	s := make([]int, 0)
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++ {
		var a int
		var m string
		fmt.Fscan(r, &m)
		if m == "push" {
			fmt.Fscan(r, &a)
		}
		switch m {
		case "push":
			s = push(s, a)
		case "pop":
			s = pop(s)
		case "size":
			size(s)
		case "empty":
			empty(s)
		case "top":
			top(s)
		}
	}
}

func push(s []int, a int) []int {
	s = append(s, a)
	return s
}

func pop(s []int) []int {
	if len(s) == 0 {
		fmt.Fprintln(w, -1)
		return s
	}
	fmt.Fprintln(w, s[len(s)-1])
	return s[:len(s)-1]
}

func size(s []int) {
	fmt.Fprintln(w, len(s))
}

func empty(s []int) {
	if len(s) == 0 {
		fmt.Fprintln(w, 1)
	} else {
		fmt.Fprintln(w, 0)
	}
}

func top(s []int) {
	if len(s) == 0 {
		fmt.Fprintln(w, -1)
		return
	}
	fmt.Fprintln(w, s[len(s)-1])
}