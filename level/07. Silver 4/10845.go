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
	var a int
	d := make([]int, 0)
	fmt.Fscanln(r, &a)
	for i := 0; i < a; i++ {
		var b int
		var c string
		fmt.Fscanf(r, "%s %d\n", &c, &b)
		switch c {
		case "push":
			d = push(d, b)
		case "pop":
			d = pop(d)
		case "size":
			size(d)
		case "empty":
			empty(d)
		case "front":
			front(d)
		case "back":
			back(d)
		}
	}
}

func push(d []int, b int) []int {
	d = append(d, b)
	return d
}
func pop(d []int) []int{
	if len(d) == 0 {
		fmt.Fprintln(w, -1)
		return nil
	} else {
		fmt.Fprintln(w, int(d[0]))
		return d[1:]
	}
}
func size(d []int) {
	fmt.Fprintln(w, len(d))
}
func empty(d []int) {
	if len(d) == 0 {
		fmt.Fprintln(w, 1)
	} else {
		fmt.Fprintln(w, 0)
	}
}
func front(d []int) {
	if len(d) == 0 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, int(d[0]))
	}
}
func back(d []int) {
	if len(d) == 0 {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, int(d[len(d)-1]))
	}
}