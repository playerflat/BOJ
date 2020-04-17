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
	var b int

	fmt.Fscanln(r, &a)
	b += color(a, "value")*10
	fmt.Fscanln(r, &a)
	b += color(a, "value")
	fmt.Fscanln(r, &a)
	b *= color(a, "multi")

	fmt.Fprintln(w, b)
}

func color(color string, mode string) int {
	switch mode {
	case "value":
		switch color {
		case "black":
			return 0
		case "brown":
			return 1
		case "red":
			return 2
		case "orange":
			return 3
		case "yellow":
			return 4
		case "green":
			return 5
		case "blue":
			return 6
		case "violet":
			return 7
		case "grey":
			return 8
		case "white":
			return 9
		}
	case "multi":
		switch color {
		case "black":
			return 1
		case "brown":
			return 10
		case "red":
			return 100
		case "orange":
			return 1000
		case "yellow":
			return 10000
		case "green":
			return 100000
		case "blue":
			return 1000000
		case "violet":
			return 10000000
		case "grey":
			return 100000000
		case "white":
			return 1000000000
		}
	}
	return 0
}